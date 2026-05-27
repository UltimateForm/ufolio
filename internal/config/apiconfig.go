package config

import (
	"log"
	"os"
)

type apiConfig struct {
	Dev              bool
	TurnstileSiteKey string
}

type secretConfig struct {
	GithubToken        string
	EdgeSignature      string
	TurnstileSecretKey string
	ResendApiKey       string
	ResendFromEmail    string
	ResendToEmail      string
}

var (
	Api    *apiConfig
	Secret *secretConfig
)

func init() {
	Api = &apiConfig{
		Dev:              os.Getenv("DEV") == "1",
		TurnstileSiteKey: os.Getenv("TURNSTILE_SITE_KEY"),
	}
	Secret = &secretConfig{
		GithubToken:        os.Getenv("GITHUB_TOKEN"),
		EdgeSignature:      os.Getenv("X_EDGE_SIGNATURE"),
		TurnstileSecretKey: os.Getenv("TURNSTILE_SECRET_KEY"),
		ResendApiKey:       os.Getenv("RESEND_API_KEY"),
		ResendFromEmail:    os.Getenv("RESEND_FROM_EMAIL"),
		ResendToEmail:      os.Getenv("RESEND_TO_EMAIL"),
	}
	if Secret.GithubToken == "" {
		log.Fatalf("GITHUB_TOKEN environment variable not set")
	}
}
