package rest

import (
	"encoding/json"
	"net/http"

	"github.com/vitorsalgado/go-location-management/internal/server/rest/router"
	"github.com/vitorsalgado/go-location-management/internal/utils/panicif"
)

// result represents health check response
type result struct {
	Status string `json:"status"`
}

// RegisterRoutes register the health check endpoint
func RegisterRoutes(r *router.RoutingMiddleware) {
	router.
		GET("/api/v1/ping").
		HandlerFunc(Ping).
		Register(r)
}

func Ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	panicif.Err(
		json.NewEncoder(w).Encode(result{"pong"}))
}
