package usecases

import (
	"encoding/json"
	"net/http"
)

// PostNewLocation is the endpoint handler to receive requests to update a user location
// Request::
// POST /api/v1/location
// Body: NewLocationRequest
// --
// Response ::
// Status: 200
func (uc *UpdateLocation) PostNewLocation(w http.ResponseWriter, r *http.Request) {
	request := NewLocationRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uc.Exec(request)

	w.WriteHeader(http.StatusOK)
}
