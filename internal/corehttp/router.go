package corehttp

import (
	"net/http"
	"strings"
)

type Router struct {
	BasePath string
	// available for compatibility, prefer Router methods over direct Mux calls
	Mux *http.ServeMux
}

func NewRouter(basePath string, mux *http.ServeMux) *Router {
	return &Router{
		BasePath: basePath,
		Mux:      mux,
	}
}

func joinPath(base, route string) string {
	if base == "" {
		return route
	}
	trimmedBase := strings.TrimRight(base, "/")
	trimmedRoute := strings.TrimLeft(route, "/")
	return trimmedBase + "/" + trimmedRoute
}

func (r *Router) Get(route string, handler http.HandlerFunc) {
	r.Mux.HandleFunc("GET "+joinPath(r.BasePath, route), handler)
}

func (r *Router) Post(route string, handler http.HandlerFunc) {
	r.Mux.HandleFunc("POST "+joinPath(r.BasePath, route), handler)
}

func (r *Router) Put(route string, handler http.HandlerFunc) {
	r.Mux.HandleFunc("PUT "+joinPath(r.BasePath, route), handler)
}

func (r *Router) Delete(route string, handler http.HandlerFunc) {
	r.Mux.HandleFunc("DELETE "+joinPath(r.BasePath, route), handler)
}
