package cmd

import (
	"net/http"
)

func addStaticRoutes(mux *http.ServeMux) {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static"))))
}
