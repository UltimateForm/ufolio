package cmd

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func RunAPI() {
	fmt.Println("Starting API server...")
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	fmt.Printf("Working directory: %s\n", wd)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseGlob("templates/*.html"))

		err := templ.ExecuteTemplate(w, "layout", map[string]string{"Arg": "worlds"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseGlob("templates/*.html"))

		err := templ.ExecuteTemplate(w, "clickResp", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
