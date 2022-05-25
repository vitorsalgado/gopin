package router

import (
	"context"
	"net/http"
)

// RoutingMiddleware holds a collection of Route and the RoutingMiddleware.
// The RoutingMiddleware handles all unhandled requests that falls to the / path.
// It will iterate through the Route collection to search for a Handler.
type RoutingMiddleware struct {
	routes []*Route
}

// Init creates a new RoutingMiddleware and sets this new instance to the http.ServeMux instance root path.
// This way, all unhandled requests (almost all) will be handled by the RoutingMiddleware.
func Init(mux *http.ServeMux) *RoutingMiddleware {
	var mdw = &RoutingMiddleware{}
	mux.Handle("/", mdw.Middleware())

	return mdw
}

// ApplyRoutesTo applies all previously configured Route[] to http.ServeMux instance
func (rm *RoutingMiddleware) ApplyRoutesTo(mux *http.ServeMux) {
	for _, r := range rm.routes {
		mux.Handle(r.path, r.handler)
	}
}

// Middleware returns the root middleware that handles all unmatched routes
func (rm *RoutingMiddleware) Middleware() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, r := range rm.routes {
			// Check first if the method matches
			if r.method == req.Method {
				// Execute the regex against the URL and extract all path params if any
				// We will iterate through the path params to create map
				if matches := r.pattern.FindAllStringSubmatch(req.URL.Path, -1); len(matches) > 0 {
					values := matches[0]
					keys := r.pattern.SubexpNames()
					params := make(map[string]string, len(values))

					for i, key := range values {
						params[keys[i]] = key
					}

					// We add the path params to a customized Context, so we can access the parameters later
					// in the Handlers
					ctx := context.WithValue(req.Context(), ParamsContextKey, params)

					// Execute the HTTP Handler with the custom Context
					r.handler.ServeHTTP(w, req.WithContext(ctx))

					return
				}
			}
		}

		http.NotFound(w, req)
	})
}
