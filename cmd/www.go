package cmd

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/UltimateForm/ufolio/internal/corehttp"
	"github.com/UltimateForm/ufolio/internal/githubapi"
)

func addWwwRoutes(router *corehttp.Router) {

	edgeSignature := os.Getenv("X_EDGE_SIGNATURE")
	// why? i want that fly.dev domain gone, outta my sight, shoo
	checkEdge := func(w http.ResponseWriter, r *http.Request) bool {
		if edgeSignature == "" || r.Header.Get("x-edge-signature") == edgeSignature {
			return true
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}

	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		log.Fatalf("GITHUB_TOKEN environment variable not set")
	}
	ghClient := githubapi.New(ghToken)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if !checkEdge(w, r) {
			return
		}
		templ := template.Must(template.ParseGlob("www/templates/*.html"))
		repos, err := ghClient.GetRepos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "layout", map[string]any{"Repos": repos})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.Get("/clicked", func(w http.ResponseWriter, r *http.Request) {
		if !checkEdge(w, r) {
			return
		}
		templ := template.Must(template.ParseGlob("www/templates/*.html"))

		err := templ.ExecuteTemplate(w, "clickResp", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

}
