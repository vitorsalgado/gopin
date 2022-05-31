package observability

import (
	"encoding/json"
	"github.com/vitorsalgado/gopin/internal/util/router"
	"net/http"

	"github.com/vitorsalgado/gopin/internal/util/panicif"
)

// Result represents health check response
type Result struct {
	Status string `json:"status"`
}

// ConfigureHealthCheck register the health check endpoint
func ConfigureHealthCheck(r *router.RoutingMiddleware) {
	router.GET("/api/v1/ping").
		HandlerFunc(Ping).
		Register(r)
}

func Ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	panicif.Err(
		json.NewEncoder(w).Encode(Result{"pong"}))
}
