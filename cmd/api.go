package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/UltimateForm/ufolio/internal/corehttp"
)

func RunAPI() {
	log.Println("Starting API server...")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	log.Printf("Working directory: %s\n", wd)

	mainMux := http.NewServeMux()

	addStaticRoutes(mainMux)

	router := corehttp.NewRouter("/", mainMux)

	addHotRoutes(corehttp.NewRouter("/hot/", router.Mux))

	addWwwRoutes(router)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", mainMux); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
