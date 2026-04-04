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

	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		log.Fatalf("GITHUB_TOKEN environment variable not set")
	}
	ghClient := githubapi.New(ghToken)
	templ := template.Must(template.ParseGlob("www/templates/*.html"))

	router.HandleRoute(corehttp.NewRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		repos, err := ghClient.GetRepos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "layout", map[string]any{"Repos": repos})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	router.HandleRoute(corehttp.NewRoute("GET", "/clicked", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseGlob("www/templates/*.html"))

		err := templ.ExecuteTemplate(w, "clickResp", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

}
