package gopin

import (
	"github.com/vitorsalgado/gopin/internal/domain"
	"github.com/vitorsalgado/gopin/internal/handlers"
	"github.com/vitorsalgado/gopin/internal/util/router"
	"github.com/vitorsalgado/gopin/internal/util/worker"
)

type LocationHandler struct {
}

// Routes register application HTTP route handlers
func Routes(
	r *router.RoutingMiddleware, dispatcher *worker.Dispatcher, repository domain.LocationRepository,
) {
	router.GET("api/v1/current_location/:user_uuid").
		HandlerFunc(handlers.NewCurrentLocationHandler(repository).Execute).
		Register(r)

	router.POST("api/v1/location").
		HandlerFunc(handlers.NewUpdateLocation(repository, dispatcher).Execute).
		Register(r)

	router.GET("api/v1/location_history/:session_uuid").
		HandlerFunc(handlers.NewHistoryHandler(repository).Execute).
		Register(r)
}
