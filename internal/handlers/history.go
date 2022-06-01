package handlers

import (
	"github.com/vitorsalgado/gopin/internal/domain"
	"github.com/vitorsalgado/gopin/internal/util/arch"
	"github.com/vitorsalgado/gopin/internal/util/check"
	"github.com/vitorsalgado/gopin/internal/util/http/httputils"
	"net/http"
)

const pSessionID = "session_uuid"

var _ arch.Handler = (*HistoryHandler)(nil)

// HistoryHandler represents the session location history listing use case
type HistoryHandler struct {
	repository domain.LocationRepository
}

// NewHistoryHandler creates a new HistoryHandler listing use case instance
func NewHistoryHandler(repository domain.LocationRepository) arch.Handler {
	return &HistoryHandler{repository}
}

// Execute is used to get a list of all location updates sent during a session.
// GET api/v1/location_history/{session_uuid}
func (uc *HistoryHandler) Execute(w http.ResponseWriter, r *http.Request) {
	params := httputils.Params(r)
	sessionId := params[pSessionID]

	if !check.IsUUIDValid(sessionId) {
		httputils.BadRequest(w, "Parameter %v is not a valid UUID", pSessionID)
		return
	}

	history, err := uc.repository.HistoryForSession(sessionId)

	if err != nil {
		httputils.Err(w, err)
		return
	}

	if history == nil || len(*history) == 0 {
		httputils.NotFound(w, "No listing found for session id %v", sessionId)
		return
	}

	httputils.OkJSON(w, history)
}
