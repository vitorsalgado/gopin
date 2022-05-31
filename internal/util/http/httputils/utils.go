package httputils

import (
	"encoding/json"
	"fmt"
	"github.com/vitorsalgado/gopin/internal/util/http/middlewares"
	"github.com/vitorsalgado/gopin/internal/util/router"
	"net/http"

	"github.com/vitorsalgado/gopin/internal/util/panicif"
)

// Headers and Content Types
const (
	HeaderContentType          = "Content-Type"
	ContentTypeApplicationJSON = "application/json"
)

// Params extract the Path Parameters, if any, contained in the http.Request Context.
func Params(r *http.Request) map[string]string {
	if values := r.Context().Value(router.ParamsContextKey); values != nil {
		return values.(map[string]string)
	}

	return nil
}

// OkJSON is a helper to return an OK (200) with a JSON body
func OkJSON(w http.ResponseWriter, model interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)

	panicif.Err(
		json.NewEncoder(w).Encode(&model))
}

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
