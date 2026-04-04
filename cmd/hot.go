package cmd

import (
	"net/http"
	"time"

	"github.com/UltimateForm/ufolio/internal/corehttp"
)

func addHotRoutes(router *corehttp.Router) {
	router.HandleRoute(corehttp.NewRoute(
		"GET",
		"/hot/live",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			ticker := time.NewTicker(time.Second * 5)
			flusher := w.(http.Flusher)
			for {
				select {
				case t := <-ticker.C:
					w.Write([]byte("data: " + t.String() + "\n\n"))
					flusher.Flush()
				case <-r.Context().Done():
					ticker.Stop()
					return
				}
			}
		},
	))

	router.HandleRoute(corehttp.NewRoute(
		"POST",
		"/hot/push",
		func(w http.ResponseWriter, r *http.Request) {

		},
	))
}
