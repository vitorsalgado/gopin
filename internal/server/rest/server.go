package rest

import (
	"net/http"

	"github.com/vitorsalgado/gopin/internal/server/rest/router"
	"github.com/vitorsalgado/gopin/internal/utils/config"
)

// Server setups an HTTP Server with basic configurations.
func Server(configurations *config.Config) (*http.ServeMux, *router.RoutingMiddleware) {
	mux := http.NewServeMux()

	// Swagger
	swagger := http.FileServer(http.Dir(configurations.SwaggerUiPath))
	mux.Handle("/docs/", http.StripPrefix("/docs/", swagger))
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "index.html") })
	mux.HandleFunc("/docs/swagger.yml", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./api/swagger.yml") })

	r := router.Init(mux)

	return mux, r
}
