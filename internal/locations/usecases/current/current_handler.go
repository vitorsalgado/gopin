package current

import (
	"net/http"

	"github.com/vitorsalgado/go-location-management/internal/server/rest"
	"github.com/vitorsalgado/go-location-management/internal/utils/validations/checks"
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
	params := rest.Params(r)
	uid := params[ParamUserUUID]

	if !checks.IsUUIDValid(uid) {
		rest.BadRequest(w, "Parameter %v is not a valid UUID", ParamUserUUID)
		return
	}

	location := uc.repository.GetCurrentLocation(uid)

	if location == nil {
		rest.NotFound(w, "We couldn't retrieve the current location of user %v", uid)
		return
	}

	rest.OkJSON(w, &location)
}
