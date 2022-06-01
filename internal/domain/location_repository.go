package domain

import (
	"database/sql"
	"time"
)

type (
	// MySQLRepository is a MySQL implementation for LocationRepository interface
	MySQLRepository struct {
		db *sql.DB
	}
)

var _ LocationRepository = (*MySQLRepository)(nil)

// NewLocationRepository returns a LocationRepository concrete implementation
func NewLocationRepository(db *sql.DB) LocationRepository {
	return &MySQLRepository{db}
}

// ReportNew reports a new location update sent by a user device
func (repository *MySQLRepository) ReportNew(location Location) error {
	stmt, err := repository.db.Prepare("INSERT INTO user_locations (user_uuid, session_uuid, lat, lng, `precision`, reported_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		location.UserID,
		location.SessionID,
		location.Latitude,
		location.Longitude,
		location.Precision,
		location.ReportedAt,
	)

	return err
}

// Current returns a user current location.
// The location is within a 10 Minutes period and with a precision of at least 1000.
func (repository *MySQLRepository) Current(id string) (*Location, error) {
	timeLimit := time.Now().Add(-10 * time.Minute)
	query := repository.db.
		QueryRow("SELECT session_uuid, lat, lng, `precision`, reported_at FROM user_locations WHERE user_uuid = ? AND reported_at >= TIME(?) AND `precision` >= 1000 ORDER BY inserted_at DESC LIMIT 1",
			id, timeLimit.String())

	var result Location
	err := query.Scan(&result.SessionID, &result.Latitude, &result.Longitude, &result.Precision, &result.ReportedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

// HistoryForSession is used to get a list of all location updates sent during a session.
func (repository *MySQLRepository) HistoryForSession(sessionId string) (*[]Location, error) {
	query, err := repository.db.Query("SELECT lat, lng, `precision`, reported_at FROM user_locations WHERE session_uuid = ?", sessionId)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	var result []Location

	for query.Next() {
		var record Location
		err = query.Scan(&record.Latitude, &record.Longitude, &record.Precision, &record.ReportedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, record)
	}

	return &result, nil
}
