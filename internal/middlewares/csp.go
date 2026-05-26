package middlewares

import "net/http"

// unsure we need this but was having some issues with cloudflare turnstile so keeping it jic
func CSP(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "script-src 'self' https://challenges.cloudflare.com https://cdn.jsdelivr.net; frame-src https://challenges.cloudflare.com; connect-src 'self'")
		next(w, r)
	})
}
