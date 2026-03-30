package cmd

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/UltimateForm/ufolio/internal/githubapi"
)

func RunAPI() {
	log.Println("Starting API server...")
	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		log.Fatalf("GITHUB_TOKEN environment variable not set")
	}
	ghClient := githubapi.New(ghToken)

	edgeSignature := os.Getenv("X_EDGE_SIGNATURE")
	// why? i want that fly.dev domain gone, outta my sight, shoo
	checkEdge := func(w http.ResponseWriter, r *http.Request) bool {
		if edgeSignature == "" || r.Header.Get("x-edge-signature") == edgeSignature {
			return true
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
		return false
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	log.Printf("Working directory: %s\n", wd)

	// i know.... will handle this later with better middleware
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !checkEdge(w, r) {
			return
		}
		templ := template.Must(template.ParseGlob("templates/*.html"))
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
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		if !checkEdge(w, r) {
			return
		}
		templ := template.Must(template.ParseGlob("templates/*.html"))

		err := templ.ExecuteTemplate(w, "clickResp", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
