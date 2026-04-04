package cmd

import (
	"io"
	"log"
	"net/http"

	"github.com/UltimateForm/ufolio/internal/corehttp"
)

func addHotRoutes(router *corehttp.Router) {
	fileChanges := make(chan string)
	defautLogger := log.Default()
	hotReloadLogger := log.New(defautLogger.Writer(), "[hot] ", defautLogger.Flags())
	router.HandleRoute(corehttp.NewRoute(
		"GET",
		"/hot/live",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			resController := http.NewResponseController(w)
			for {
				select {
				case f := <-fileChanges:
					// ideally we would be doing specific logic per file change
					// for example, depending on the file type, reloading the page or sending a specific event
					// but for now, we just send a generic "file change detected" message
					// so all changes will trigger a reload
					hotReloadLogger.Printf("file change detected: %s, reloading\n", f)
					w.Write([]byte("data: RELOAD\n\n"))
					if err := resController.Flush(); err != nil {
						hotReloadLogger.Printf("flush error: %v\n", err)
						return
					}
				case <-r.Context().Done():
					return
				}
			}
		},
	))

	router.HandleRoute(corehttp.NewRoute(
		"POST",
		"/hot/push",
		func(w http.ResponseWriter, r *http.Request) {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fileChanges <- string(bodyBytes)
		},
	))
}
