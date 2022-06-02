package router

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type contextKey int

const (
	backSlash                   = "/"
	pathParamPrefix             = ":"
	ParamsContextKey contextKey = iota
)

type (
	// Route represents a route in the application.
	// It holds values that facilitate handling and dispatching requests to the http.Handler specified in the
	// handler property.
	Route struct {
		path    string
		method  string
		pattern *regexp.Regexp
		handler http.Handler
	}

	// RouteBuilder is utility to facilitate the registration of a Route in the application.
	// The preferred way to begin with this Builder is using: GET() or POST(). (Just these two methods for now).
	// After setting the HTTP Method and Handler or HandlerFunc, call Register to finish registration.
	RouteBuilder struct {
		path    string
		method  string
		handler http.Handler
	}
)

// GET init a RouteBuilder with HTTP Method set as GET
// The path must be in the format:
func GET(path string) *RouteBuilder {
	return &RouteBuilder{method: http.MethodGet, path: path}
}

// POST init a RouteBuilder with HTTP Method set as POST
// The path must be in the format:
func POST(path string) *RouteBuilder {
	return &RouteBuilder{method: http.MethodPost, path: path}
}

// Handler defines a http.Handler for the Route
func (b *RouteBuilder) Handler(handler http.Handler) *RouteBuilder {
	b.handler = handler
	return b
}

// HandlerFunc defines a func(http.ResponseWriter, *http.Request) for the Route
func (b *RouteBuilder) HandlerFunc(handler func(http.ResponseWriter, *http.Request)) *RouteBuilder {
	b.handler = http.HandlerFunc(handler)
	return b
}

// Register is the last func to be called after setting all Route configurations.
// This method process the Route and all possible path params defined.
func (b *RouteBuilder) Register(router *RoutingMiddleware) {
	var parts = strings.Split(b.path, backSlash)
	var sb strings.Builder

	for i, p := range parts {
		if i > 0 {
			sb.WriteString(backSlash)
		}

		if hasPathParam(p) {
			sb.WriteString(paramRegex(extractPathParamName(p)))
			continue
		}

		sb.WriteString(p)
	}

	if strings.HasSuffix(b.path, backSlash) {
		sb.WriteString(backSlash)
	}

	url := sb.String()
	regex := regexp.MustCompile(url)
	r := &Route{path: b.path, method: b.method, pattern: regex, handler: b.handler}

	router.routes = append(router.routes, r)
}

func hasPathParam(p string) bool {
	return strings.HasPrefix(p, pathParamPrefix)
}

func extractPathParamName(p string) string {
	return p[1:]
}

func paramRegex(p string) string {
	return fmt.Sprintf("(?P<%v>\\w.+)", p)
}
