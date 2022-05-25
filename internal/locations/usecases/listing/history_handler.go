package listing

import (
	"net/http"

	"github.com/vitorsalgado/go-location-management/internal/server/rest"
	"github.com/vitorsalgado/go-location-management/internal/utils/validations/checks"
)

const ParamSessionUUID = "session_uuid"

// GetSessionLocationHistory is used to get a list of all location updates sent during a session.
// Request::
// GET api/v1/location_history/{session_uuid}
// --
// Response:
// Status: 200
// Body: SessionLocation[]
func (uc *History) GetSessionLocationHistory(w http.ResponseWriter, r *http.Request) {
	params := rest.Params(r)
	sessionId := params[ParamSessionUUID]

	if !checks.IsUUIDValid(sessionId) {
		rest.BadRequest(w, "Parameter %v is not a valid UUID", ParamSessionUUID)
		return
	}

	history := uc.repository.HistoryFor(sessionId)

	if history == nil || len(history) == 0 {
		rest.NotFound(w, "No listing found for session id %v", sessionId)
		return
	}

	rest.OkJSON(w, history)
}
