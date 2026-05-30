package middlewares

import "net/http"

func Cache(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Cache-Control", "public, max-age=86400")
		}
		next(w, r)
	})
}
