package cmd

import (
	"log"
	"net/http"

	"github.com/UltimateForm/ufolio/internal/config"
	"github.com/UltimateForm/ufolio/internal/corehttp"
	"github.com/UltimateForm/ufolio/internal/middlewares"
)

func RunAPI() {
	log.Println("Starting API server...")

	if config.Api.Dev {
		log.Println("RUNNING IN DEV MODE")
	}

	// i want to be able to change this without caring about cmd handlers
	// so i am abstracting it into a separate internal package
	router := corehttp.NewRouter("/").With(middlewares.EnforceEdge).With(middlewares.Logging)

	addStaticRoutes(router)

	if config.Api.Dev {
		addHotRoutes(router)
	}

	addWwwRoutes(router)

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
