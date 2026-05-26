package cmd

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/UltimateForm/ufolio/internal/turnstile"
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
		log.Print("subject or body is empty")
		templ.ExecuteTemplate(w, "contact-me-notok", nil)
		return
	}
	turnstileToken := r.FormValue("cf-turnstile-response")
	if turnstileToken == "" {
		log.Printf("unable to verify form submission, no cf-turnstile-response token, or empty")
		templ.ExecuteTemplate(w, "contact-me-notok", nil)
		return
	}
	ctx, _ := context.WithTimeout(r.Context(), time.Second*20)
	if err := turnstile.ValidateToken(ctx, turnstileToken); err != nil {
		log.Printf("contact-me form submission failed verification: %v", err)
		templ.ExecuteTemplate(w, "contact-me-notok", nil)
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
