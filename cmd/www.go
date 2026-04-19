package cmd

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/UltimateForm/ufolio/internal/config"
	"github.com/UltimateForm/ufolio/internal/corehttp"
	"github.com/UltimateForm/ufolio/internal/githubapi"
)

func flatMapKeys(m map[string]any) []string {
	keys := make([]string, 0)
	for k, v := range m {
		if _, ok := v.(map[string]any); ok {
			keys = append(keys, flatMapKeys(v.(map[string]any))...)
		}
		keys = append(keys, k)
	}
	return keys
}

func loadTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"link": func(url, text string) template.HTML {
			return template.HTML("<a href=\"" + url + "\" target=\"_blank\" rel=\"noopener noreferrer\">" + text + "</a>")
		},
	}
}

func loadMainTemplates() *template.Template {
	rootTemplateFileNames, _ := filepath.Glob("www/templates/*.html")
	return template.Must(template.New("main").Funcs(loadTemplateFuncs()).ParseFiles(rootTemplateFileNames...))
}

func loadSkillsTemplates() *template.Template {
	skillsTemplateFileNames, _ := filepath.Glob("www/templates/skills/*.html")
	return template.Must(template.New("skills").Funcs(loadTemplateFuncs()).ParseFiles(skillsTemplateFileNames...))
}

func getSkillsHtml(templ *template.Template, techKeys []string) string {
	// todo: consider handling errors on these
	var skillsHtml strings.Builder
	for _, key := range techKeys {
		err := templ.ExecuteTemplate(&skillsHtml, key, nil)
		if err != nil {
			log.Printf("could not render skill template for %s: %v", key, err)
		}
	}
	return skillsHtml.String()
}

func addWwwRoutes(router *corehttp.Router) {

	ghClient := githubapi.New(config.Api.GithubToken)

	mainTemplates := loadMainTemplates()
	skillsTemplate := loadSkillsTemplates()

	techJson, err := os.ReadFile("www/static/tech-tree.json")
	if err != nil {
		log.Fatal(err)
	}
	techData := make(map[string]any)
	err = json.Unmarshal(techJson, &techData)
	if err != nil {
		log.Fatal(err)
	}
	techKeys := flatMapKeys(techData)
	skillsHtml := getSkillsHtml(skillsTemplate, techKeys)

	router.HandleRoute(corehttp.NewRoute("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		if config.Api.Dev {
			// reload templates on every request in dev mode
			mainTemplates = loadMainTemplates()
			skillsTemplate = loadSkillsTemplates()
			skillsHtml = getSkillsHtml(skillsTemplate, techKeys)
		}
		repos, err := ghClient.GetRepos(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = mainTemplates.ExecuteTemplate(w, "layout", map[string]any{"Repos": repos, "Dev": config.Api.Dev, "Tech": techData, "TechSkillsHtml": template.HTML(skillsHtml)})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	router.HandleRoute(corehttp.NewRoute("POST", "/clicked", func(w http.ResponseWriter, r *http.Request) {
		if config.Api.Dev {
			// reload templates on every request in dev mode
			mainTemplates = template.Must(template.ParseGlob("www/templates/*.html"))
		}
		err := mainTemplates.ExecuteTemplate(w, "clickResp", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

}
