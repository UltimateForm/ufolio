package cmd

import (
	"html/template"
	"log"
	"net/http"
)

func handleContact(w http.ResponseWriter, r *http.Request, templ *template.Template) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	subject := r.FormValue("subject")
	body := r.FormValue("body")
	if subject == "" || body == "" {
		http.Error(w, "Subject and body are required", http.StatusBadRequest)
		return
	}
	log.Printf("#### contact me form submission:\nsubject: %v\nbody: %v", subject, body)
	templ.ExecuteTemplate(w, "contact-me-ok", nil)
}

func createHandleContact(templ *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleContact(w, r, templ)
	}
}
