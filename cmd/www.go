package cmd

import (
	"html/template"
	"net/http"

	"github.com/UltimateForm/ufolio/internal/config"
	"github.com/UltimateForm/ufolio/internal/corehttp"
	"github.com/UltimateForm/ufolio/internal/githubapi"
)

func addWwwRoutes(router *corehttp.Router) {

	ghClient := githubapi.New(config.Api.GithubToken)

	templ := template.Must(template.ParseGlob("www/templates/*.html"))

	router.HandleRoute(corehttp.NewRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		if config.Api.Dev {
			// reload templates on every request in dev mode
			templ = template.Must(template.ParseGlob("www/templates/*.html"))
		}
		repos, err := ghClient.GetRepos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "layout", map[string]any{"Repos": repos, "Dev": config.Api.Dev})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	router.HandleRoute(corehttp.NewRoute("POST", "/clicked", func(w http.ResponseWriter, r *http.Request) {
		if config.Api.Dev {
			// reload templates on every request in dev mode
			templ = template.Must(template.ParseGlob("www/templates/*.html"))
		}
		err := templ.ExecuteTemplate(w, "clickResp", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

}
