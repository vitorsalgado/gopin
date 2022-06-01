package handlers

import (
	"github.com/vitorsalgado/gopin/internal/domain"
	"github.com/vitorsalgado/gopin/internal/util/arch"
	"github.com/vitorsalgado/gopin/internal/util/check"
	"github.com/vitorsalgado/gopin/internal/util/http/httputils"
	"net/http"
)

const pUserID = "user_uuid"

var _ arch.Handler = (*CurrentLocationHandler)(nil)

// CurrentLocationHandler represents the current location use case
type CurrentLocationHandler struct {
	repository domain.LocationRepository
}

// NewCurrentLocationHandler creates current location use case instance
func NewCurrentLocationHandler(repository domain.LocationRepository) arch.Handler {
	return &CurrentLocationHandler{repository: repository}
}

// Execute is the handler for requesting a current location
// GET api/v1/current_location/:user_uuid
func (uc *CurrentLocationHandler) Execute(w http.ResponseWriter, r *http.Request) {
	params := httputils.Params(r)
	uid := params[pUserID]

	if !check.IsUUIDValid(uid) {
		httputils.BadRequest(w, "Parameter %v is not a valid UUID", pUserID)
		return
	}

	location, err := uc.repository.Current(uid)

	if err != nil {
		httputils.Err(w, err)
		return
	}

	if location == nil {
		httputils.NotFound(w, "We couldn't retrieve the current location of user %v", uid)
		return
	}

	httputils.OkJSON(w, &location)
}
