package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vitorsalgado/go-location-management/internal/server/rest/middlewares"
	"github.com/vitorsalgado/go-location-management/internal/server/rest/router"
	"github.com/vitorsalgado/go-location-management/internal/utils/panicif"
)

// Headers and Content Types
const (
	HeaderContentType          = "Content-Type"
	ContentTypeApplicationJSON = "application/json"
)

// Parameter Utilities
// --

// Params extract the Path Parameters, if any, contained in the http.Request Context.
func Params(r *http.Request) map[string]string {
	if values := r.Context().Value(router.ParamsContextKey); values != nil {
		return values.(map[string]string)
	}

	return nil
}

// Success Utilities
// --

// OkJSON is a helper to return an OK (200) with a JSON body
func OkJSON(w http.ResponseWriter, model interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)

	panicif.Err(
		json.NewEncoder(w).Encode(&model))
}

// Error Utilities
// --

// BadRequest creates a Bad Request (400) response
func BadRequest(w http.ResponseWriter, format string, a ...interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusBadRequest)

	panicif.Err(
		json.NewEncoder(w).Encode(&middlewares.ApiError{Message: fmt.Sprintf(format, a...)}))
}

// NotFound creates a Not Found (404) response
func NotFound(w http.ResponseWriter, format string, a ...interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusNotFound)

	panicif.Err(
		json.NewEncoder(w).Encode(&middlewares.ApiError{Message: fmt.Sprintf(format, a...)}))
}
