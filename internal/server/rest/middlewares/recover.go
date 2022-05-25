package middlewares

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/vitorsalgado/gopin/internal/utils/panicif"
)

// ApiError is the default response model for all failed API calls.
// Code: should be a business code. A code that may trigger some specific behaviour on a client caller.
// Message: short descriptive error message.
// DetailedMessage: should contain a more technical message.
type ApiError struct {
	Code            string `json:"code,omitempty"`
	Message         string `json:"message"`
	DetailedMessage string `json:"detailed_message,omitempty"`
}

// Recovery is a recover() middleware for all unhandled errors
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recovery := recover(); recovery != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				var msg string
				switch x := recovery.(type) {
				case string:
					msg = x
				case error:
					msg = x.Error()
				default:
					msg = http.StatusText(http.StatusInternalServerError)
				}

				log.Error().
					Timestamp().
					Stack().
					Interface("recover", recovery).
					Msg(msg)

				panicif.Err(
					json.NewEncoder(w).Encode(&ApiError{Message: msg}))
			}

		}()

		next.ServeHTTP(w, r)
	})
}
