package corehttp

import (
	"net/http"
	"strings"
)

type Route struct {
	// could be a enum but opting for string to keep it simple
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func (src *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	src.HandlerFunc(w, r)
}

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func (src *Route) With(middleware Middleware) *Route {
	src.HandlerFunc = middleware(src.HandlerFunc)
	return src
}

func NewRoute(method, path string, handlerFunc http.HandlerFunc) *Route {
	return &Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
	}
}

type Router struct {
	Route
	subroutes map[string]*Route
	BasePath  string
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.HandlerFunc(w, req)
}

func NewRouter(basePath string) *Router {
	router := &Router{
		Route: Route{
			HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
				return
			},
		},
		subroutes: make(map[string]*Route),
		BasePath:  basePath,
	}
	router.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		normalized := normalizeRouteSlug(r.Method, router.BasePath, r.URL.Path)

		if route, ok := router.subroutes[normalized]; ok {
			route.ServeHTTP(w, r)
			return
		}

		// all of the wildcard matching here is just for a single wildcard route we have (/static/*)
		// this is not a public lib so we could just hardcode it here, ie if r.Method == "GET" && strings.HasPrefix(r.URL.Path, "/static/") {}
		// but i wanna try this out just for fun

		// iterate over subroutes and find a wildcard match
		for _, route := range router.subroutes {
			if !strings.HasSuffix(route.Path, "*") {
				continue
			}
			if route.Method != r.Method {
				continue
			}
			if strings.HasPrefix(r.URL.Path, strings.TrimRight(route.Path, "*")) {
				route.ServeHTTP(w, r)
				return
			}
		}

		http.NotFound(w, r)
	}
	return router
}

func joinPath(base, route string) string {
	if base == "" {
		return route
	}
	trimmedBase := strings.TrimRight(base, "/")
	trimmedRoute := strings.TrimLeft(route, "/")
	return trimmedBase + "/" + trimmedRoute
}

func normalizeRouteSlug(method string, basePath string, path string) string {
	return method + " " + joinPath(basePath, path)
}

func (r *Router) With(middleware Middleware) *Router {
	r.HandlerFunc = middleware(r.HandlerFunc)
	return r
}

func (r *Router) HandleRoute(route *Route) {
	r.subroutes[normalizeRouteSlug(route.Method, r.BasePath, route.Path)] = route
}
