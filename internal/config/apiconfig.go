package config

import (
	"log"
	"os"
)

type apiConfig struct {
	GithubToken        string
	EdgeSignature      string
	TurnstileSiteKey   string
	TurnstileSecretKey string
	Dev                bool
}

var Api *apiConfig

func init() {
	Api = &apiConfig{
		GithubToken:        os.Getenv("GITHUB_TOKEN"),
		EdgeSignature:      os.Getenv("X_EDGE_SIGNATURE"),
		Dev:                os.Getenv("DEV") == "1",
		TurnstileSiteKey:   os.Getenv("TURNSTILE_SITE_KEY"),
		TurnstileSecretKey: os.Getenv("TURNSTILE_SECRET_KEY"),
	}
	if Api.GithubToken == "" {
		log.Fatalf("GITHUB_TOKEN environment variable not set")
	}
}
