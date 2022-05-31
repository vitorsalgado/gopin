package usecases

import (
	"github.com/vitorsalgado/gopin/internal/util/check"
	"github.com/vitorsalgado/gopin/internal/util/http/httputils"
	"net/http"
)

const paramSessionUUID = "session_uuid"

// GetSessionLocationHistory is used to get a list of all location updates sent during a session.
// Request::
// GET api/v1/location_history/{session_uuid}
// --
// Response:
// Status: 200
// Body: SessionLocation[]
func (uc *History) GetSessionLocationHistory(w http.ResponseWriter, r *http.Request) {
	params := httputils.Params(r)
	sessionId := params[paramSessionUUID]

	if !check.IsUUIDValid(sessionId) {
		httputils.BadRequest(w, "Parameter %v is not a valid UUID", paramSessionUUID)
		return
	}

	history := uc.repository.HistoryForSession(sessionId)

	if history == nil || len(history) == 0 {
		httputils.NotFound(w, "No listing found for session id %v", sessionId)
		return
	}

	httputils.OkJSON(w, history)
}
