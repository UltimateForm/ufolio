package cmd

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/UltimateForm/ufolio/internal/mail"
	"github.com/UltimateForm/ufolio/internal/turnstile"
)

var contactLogger = log.New(log.Default().Writer(), "[contact] ", log.Default().Flags())

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
		contactLogger.Print("subject or body is empty")
		templ.ExecuteTemplate(w, "contact-me-notok", nil)
		return
	}
	turnstileToken := r.FormValue("cf-turnstile-response")
	if turnstileToken == "" {
		contactLogger.Printf("unable to verify form submission, no cf-turnstile-response token, or empty")
		templ.ExecuteTemplate(w, "contact-me-notok", nil)
		return
	}
	ctx, _ := context.WithTimeout(r.Context(), time.Second*20)
	if err := turnstile.ValidateToken(ctx, turnstileToken); err != nil {
		contactLogger.Printf("contact-me form submission failed verification: %v", err)
		templ.ExecuteTemplate(w, "contact-me-notok", nil)
		return
	}
	email := r.FormValue("email")
	emailCtx, _ := context.WithTimeout(r.Context(), time.Second*30)
	err := mail.SendEmail(emailCtx, email, subject, body)
	if err != nil {
		contactLogger.Printf("failed to send email: %v", err)
		templ.ExecuteTemplate(w, "contact-me-notok", nil)
		return
	}
	templ.ExecuteTemplate(w, "contact-me-ok", nil)
}

func createHandleContact(templ *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleContact(w, r, templ)
	}
}
