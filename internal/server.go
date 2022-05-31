package gopin

import (
	"github.com/vitorsalgado/gopin/internal/config"
	"github.com/vitorsalgado/gopin/internal/util/router"
	"net/http"
)

// Server setups an HTTP Server with basic configurations.
func Server(configurations *config.Config) (*http.ServeMux, *router.RoutingMiddleware) {
	mux := http.NewServeMux()

	swagger := http.FileServer(http.Dir(configurations.SwaggerUiPath))

	mux.Handle("/docs/", http.StripPrefix("/docs/", swagger))
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/docs/", http.StatusTemporaryRedirect)
	})
	mux.HandleFunc("/docs/swagger.yml", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./docs/openapi/swagger.yml") })

	r := router.Init(mux)

	return mux, r
}
