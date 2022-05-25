package locations

import (
	"github.com/vitorsalgado/go-location-management/internal/locations/usecases/listing"

	"github.com/vitorsalgado/go-location-management/internal/locations/domain"
	"github.com/vitorsalgado/go-location-management/internal/locations/usecases/adding"
	"github.com/vitorsalgado/go-location-management/internal/locations/usecases/current"
	"github.com/vitorsalgado/go-location-management/internal/server/rest/router"
	"github.com/vitorsalgado/go-location-management/internal/utils/worker"
)

// RegisterRoutes location routes
func RegisterRoutes(r *router.RoutingMiddleware, dispatcher *worker.Dispatcher, repository domain.Repository) {
	cur := current.NewCurrentLocation(repository)
	history := listing.NewHistory(repository)
	addLocation := adding.NewUpdateLocation(repository, dispatcher)

	router.
		GET("api/v1/current_location/:user_uuid").
		HandlerFunc(cur.GetCurrentLocation).
		Register(r)

	router.
		POST("api/v1/location").
		HandlerFunc(addLocation.PostNewLocation).
		Register(r)

	router.
		GET("api/v1/location_history/:session_uuid").
		HandlerFunc(history.GetSessionLocationHistory).
		Register(r)
}
