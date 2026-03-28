package cmd

import (
	"html/template"
	"net/http"
)

func RunAPI() {

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
	http.ListenAndServe(":3000", nil)
}
