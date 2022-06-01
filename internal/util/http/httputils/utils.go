package httputils

import (
	"encoding/json"
	"fmt"
	"github.com/vitorsalgado/gopin/internal/util/http/middlewares"
	"github.com/vitorsalgado/gopin/internal/util/router"
	"net/http"
)

// Headers and Content Types
const (
	HeaderContentType          = "content-type"
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
func OkJSON[T interface{}](w http.ResponseWriter, msg *T) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)

	respond(w, &msg)
}

// BadRequest creates a Bad Request (400) response
func BadRequest(w http.ResponseWriter, msg string, a ...interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusBadRequest)

	respond(w, &middlewares.ApiError{Message: fmt.Sprintf(msg, a...)})
}

// NotFound creates a Not Found (404) response
func NotFound(w http.ResponseWriter, msg string, a ...interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusNotFound)

	respond(w, &middlewares.ApiError{Message: fmt.Sprintf(msg, a...)})
}

func Err(w http.ResponseWriter, err error) {
	w.Header().Set(HeaderContentType, ContentTypeApplicationJSON)
	w.WriteHeader(http.StatusNotFound)

	_ = json.NewEncoder(w).Encode(&middlewares.ApiError{Message: err.Error()})
}

func respond[M any](w http.ResponseWriter, model *M) {
	err := json.NewEncoder(w).Encode(model)

	if err != nil {
		Err(w, err)
	}
}
