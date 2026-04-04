package middlewares

import (
	"net/http"
	"os"
)

func EnforceEdge(next http.HandlerFunc) http.HandlerFunc {
	edgeSignature := os.Getenv("X_EDGE_SIGNATURE")
	// why? i want that fly.dev domain gone, outta my sight, shoo
	return func(w http.ResponseWriter, r *http.Request) {
		if edgeSignature == "" || r.Header.Get("x-edge-signature") == edgeSignature {
			next(w, r)
			return
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}
