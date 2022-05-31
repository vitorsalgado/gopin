package usecases

import (
	"github.com/vitorsalgado/gopin/internal/core"
	"github.com/vitorsalgado/gopin/internal/util/router"
	"github.com/vitorsalgado/gopin/internal/util/worker"
)

type LocationHandler struct {
}

// RegisterLocationRoutes location routes
func RegisterLocationRoutes(
	r *router.RoutingMiddleware, dispatcher *worker.Dispatcher, repository core.LocationRepository,
) {
	cur := NewCurrentLocation(repository)
	history := NewHistory(repository)
	addLocation := NewUpdateLocation(repository, dispatcher)

	router.GET("api/v1/current_location/:user_uuid").
		HandlerFunc(cur.GetCurrentLocation).
		Register(r)

	router.POST("api/v1/location").
		HandlerFunc(addLocation.PostNewLocation).
		Register(r)

	router.GET("api/v1/location_history/:session_uuid").
		HandlerFunc(history.GetSessionLocationHistory).
		Register(r)
}
