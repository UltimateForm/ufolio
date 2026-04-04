package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/UltimateForm/ufolio/internal/corehttp"
	"github.com/UltimateForm/ufolio/internal/middlewares"
)

func RunAPI() {
	log.Println("Starting API server...")

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	log.Printf("Working directory: %s\n", wd)

	// i want to be able to change this without caring about cmd handlers
	// so i am abstracting it into a separate internal package
	router := corehttp.NewRouter("/").With(middlewares.EnforceEdge).With(middlewares.Logging)

	addStaticRoutes(router)

	addHotRoutes(router)

	addWwwRoutes(router)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
