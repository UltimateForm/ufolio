package githubapi

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type GhClient struct {
	client      *http.Client
	bearerToken string
	baseUrl     string
	// switch for slog if extended in the future
	logger    *log.Logger
	userAgent string
}

func New(bearerToken string) *GhClient {
	logger := log.New(log.Default().Writer(), "[GhClient] ", log.Default().Flags())
	return &GhClient{
		client:      http.DefaultClient,
		bearerToken: bearerToken,
		baseUrl:     "https://api.github.com",
		// this is not compliant, fix later
		userAgent: "github.com/UltimateForm/ufolio",
		logger:    logger,
	}
}

func (src *GhClient) GetRepos(ctx context.Context) ([]Repo, error) {
	src.logger.Printf("executing GetRepos")
	req, err := http.NewRequestWithContext(ctx, "GET", src.baseUrl+"/user/repos?sort=pushed&per_page=100&affiliation=owner", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+src.bearerToken)
	req.Header.Add("User-Agent", src.userAgent)
	resp, err := src.client.Do(req)
	if err != nil {
		src.logger.Printf("error executing GetRepos: %v", err)
		return nil, err
	}
	src.logger.Printf("GetRepos response status: %s", resp.Status)
	bodyBytes, err := io.ReadAll(resp.Body)
	src.logger.Printf("GetRepos response body len: %d", len(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var repos []Repo
	err = json.Unmarshal(bodyBytes, &repos)
	if err != nil {
		src.logger.Printf("error unmarshalling GetRepos response: %v", err)
		return nil, err
	}
	src.logger.Printf("GetRepos response repos: %d repos", len(repos))
	return repos, nil
}
