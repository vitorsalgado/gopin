package locations

import (
	"database/sql"
	"time"

	"github.com/vitorsalgado/go-location-management/internal/locations/domain"
	"github.com/vitorsalgado/go-location-management/internal/utils/panicif"
)

type (
	// MySQLRepository is a MySQL implementation for Repository interface
	MySQLRepository struct {
		db *sql.DB
	}
)

// NewRepository returns a Repository concrete implementation
func NewRepository(db *sql.DB) domain.Repository {
	return &MySQLRepository{db}
}

// ReportNewLocation reports a new location update sent by a user device
func (repository *MySQLRepository) ReportNewLocation(location domain.Location) {
	stmt, err := repository.db.Prepare("INSERT INTO user_locations (user_uuid, session_uuid, lat, lng, `precision`, reported_at) VALUES (?, ?, ?, ?, ?, ?)")
	panicif.Err(err)

	defer func() { panicif.Err(stmt.Close()) }()

	_, err = stmt.Exec(
		location.UserID,
		location.SessionID,
		location.Latitude,
		location.Longitude,
		location.Precision,
		location.ReportedAt,
	)
	panicif.Err(err)
}

// GetCurrentLocation returns a user current location.
// The location is within a 10 Minutes period and with a precision of at least 1000.
func (repository *MySQLRepository) GetCurrentLocation(id string) *domain.Location {
	timeLimit := time.Now().Add(-10 * time.Minute)
	query := repository.db.
		QueryRow("SELECT session_uuid, lat, lng, `precision`, reported_at FROM user_locations WHERE user_uuid = ? AND reported_at >= TIME(?) AND `precision` >= 1000 ORDER BY inserted_at DESC LIMIT 1",
			id, timeLimit.String())

	var result domain.Location
	err := query.Scan(&result.SessionID, &result.Latitude, &result.Longitude, &result.Precision, &result.ReportedAt)

	if err == sql.ErrNoRows {
		return nil
	}

	panicif.Err(err)

	return &result
}

// HistoryFor is used to get a list of all location updates sent during a session.
func (repository *MySQLRepository) HistoryFor(sessionId string) []domain.Location {
	query, err := repository.db.Query("SELECT lat, lng, `precision`, reported_at FROM user_locations WHERE session_uuid = ?", sessionId)
	panicif.Err(err)

	defer func() { panicif.Err(query.Close()) }()

	var result []domain.Location

	for query.Next() {
		var record domain.Location
		err = query.Scan(&record.Latitude, &record.Longitude, &record.Precision, &record.ReportedAt)
		panicif.Err(err)
		result = append(result, record)
	}

	return result
}
