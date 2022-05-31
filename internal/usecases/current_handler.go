package usecases

import (
	"github.com/vitorsalgado/gopin/internal/util/check"
	"github.com/vitorsalgado/gopin/internal/util/http/httputils"
	"net/http"
)

const ParamUserUUID = "user_uuid"

// GetCurrentLocation is the handler for requesting a current location
// Request::
// GET api/v1/current_location/:user_uuid
// --
// Response::
// Status: 200
// Body: CurrentLocation[]
func (uc *Location) GetCurrentLocation(w http.ResponseWriter, r *http.Request) {
	params := httputils.Params(r)
	uid := params[ParamUserUUID]

	if !check.IsUUIDValid(uid) {
		httputils.BadRequest(w, "Parameter %v is not a valid UUID", ParamUserUUID)
		return
	}

	location := uc.repository.GetCurrent(uid)

	if location == nil {
		httputils.NotFound(w, "We couldn't retrieve the current location of user %v", uid)
		return
	}

	httputils.OkJSON(w, &location)
}
