package adding

import (
	"time"

	"github.com/vitorsalgado/gopin/internal/locations/domain"
	"github.com/vitorsalgado/gopin/internal/utils/worker"
)

type (
	// NewLocationRequest represents a new location update request
	NewLocationRequest struct {
		UserID     string    `json:"user_uuid"`
		SessionID  string    `json:"session_uuid"`
		Latitude   float64   `json:"lat"`
		Longitude  float64   `json:"lng"`
		Precision  float64   `json:"precision"`
		ReportedAt time.Time `json:"reported_at"`
	}

	// UpdateLocation represents location update use case
	UpdateLocation struct {
		repository domain.Repository
		dispatcher *worker.Dispatcher
	}
)

const JobReportLocationUpdate = "Location Update"

// NewUpdateLocation creates new location update use case instance
func NewUpdateLocation(repository domain.Repository, dispatcher *worker.Dispatcher) *UpdateLocation {
	return &UpdateLocation{repository, dispatcher}
}

// Exec executes the UpdateLocation UseCase
func (uc *UpdateLocation) Exec(args NewLocationRequest) {
	uc.dispatcher.
		Dispatch(
			func() (id string, err error) {

				uc.repository.ReportNewLocation(domain.Location{
					UserID:     args.UserID,
					SessionID:  args.SessionID,
					Latitude:   args.Latitude,
					Longitude:  args.Longitude,
					Precision:  args.Precision,
					ReportedAt: args.ReportedAt,
				})

				return JobReportLocationUpdate, nil
			})
}
