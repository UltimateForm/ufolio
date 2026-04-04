package cmd

import (
	"net/http"

	"github.com/UltimateForm/ufolio/internal/corehttp"
)

func addStaticRoutes(router *corehttp.Router) {
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("www/static")))
	router.HandleRoute(corehttp.NewRoute(
		"GET",
		"/static/*",
		func(w http.ResponseWriter, r *http.Request) {
			fileServer.ServeHTTP(w, r)
		},
	))
}
