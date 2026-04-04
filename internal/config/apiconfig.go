package config

import (
	"log"
	"os"
)

type apiConfig struct {
	GithubToken   string
	EdgeSignature string
	Dev           bool
}

var Api *apiConfig

func init() {
	Api = &apiConfig{
		GithubToken:   os.Getenv("GITHUB_TOKEN"),
		EdgeSignature: os.Getenv("X_EDGE_SIGNATURE"),
		Dev:           os.Getenv("DEV") == "1",
	}
	if Api.GithubToken == "" {
		log.Fatalf("GITHUB_TOKEN environment variable not set")
	}
}
