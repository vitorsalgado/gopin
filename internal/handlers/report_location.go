package handlers

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/vitorsalgado/gopin/internal/domain"
	"github.com/vitorsalgado/gopin/internal/util/arch"
	"github.com/vitorsalgado/gopin/internal/util/http/httputils"
	"net/http"
	"time"

	"github.com/vitorsalgado/gopin/internal/util/worker"
)

const jobReportLocation = "ReportLocation"

var _ arch.Handler = (*ReportLocationHandler)(nil)

type (
	// ReportLocationRequest represents a new location update request
	ReportLocationRequest struct {
		UserID     string    `json:"user_uuid"`
		SessionID  string    `json:"session_uuid"`
		Latitude   float64   `json:"lat"`
		Longitude  float64   `json:"lng"`
		Precision  float64   `json:"precision"`
		ReportedAt time.Time `json:"reported_at"`
	}

	// ReportLocationHandler represents location update use case
	ReportLocationHandler struct {
		repository domain.LocationRepository
		dispatcher *worker.Dispatcher
	}
)

// NewReportLocationHandler creates new location update use case instance
func NewReportLocationHandler(
	repository domain.LocationRepository, dispatcher *worker.Dispatcher,
) arch.Handler {
	return &ReportLocationHandler{repository, dispatcher}
}

// Execute is the endpoint handler to receive requests to update a user location
// POST /api/v1/location
func (uc *ReportLocationHandler) Execute(w http.ResponseWriter, r *http.Request) {
	request := ReportLocationRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		httputils.Err(w, err)
		return
	}

	uc.dispatcher.Dispatch(func() (id string, err error) {
		err = uc.repository.ReportNew(domain.Location{
			UserID:     request.UserID,
			SessionID:  request.SessionID,
			Latitude:   request.Latitude,
			Longitude:  request.Longitude,
			Precision:  request.Precision,
			ReportedAt: request.ReportedAt,
		})

		if err != nil {
			log.Warn().Err(err).
				Str("user_id", request.UserID).
				Str("session_id", request.SessionID).
				Float64("lat", request.Latitude).
				Float64("lng", request.Longitude).
				Msgf("error report new location for user %s", request.UserID)
		}

		return jobReportLocation, err
	})

	w.WriteHeader(http.StatusAccepted)
}
