package usecases

import (
	"github.com/vitorsalgado/gopin/internal/core"
	"time"

	"github.com/vitorsalgado/gopin/internal/util/worker"
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
		repository core.LocationRepository
		dispatcher *worker.Dispatcher
	}
)

const JobReportLocationUpdate = "Location Update"

// NewUpdateLocation creates new location update use case instance
func NewUpdateLocation(repository core.LocationRepository, dispatcher *worker.Dispatcher) *UpdateLocation {
	return &UpdateLocation{repository, dispatcher}
}

// Exec executes the UpdateLocation UseCase
func (uc *UpdateLocation) Exec(args NewLocationRequest) {
	uc.dispatcher.
		Dispatch(
			func() (id string, err error) {

				uc.repository.ReportNew(core.Location{
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
