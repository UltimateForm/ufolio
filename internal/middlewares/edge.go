package middlewares

import (
	"net/http"

	"github.com/UltimateForm/ufolio/internal/config"
)

func EnforceEdge(next http.HandlerFunc) http.HandlerFunc {
	// why? i want that fly.dev domain gone, outta my sight, shoo
	return func(w http.ResponseWriter, r *http.Request) {
		if config.Secret.EdgeSignature == "" || r.Header.Get("x-edge-signature") == config.Secret.EdgeSignature {
			next(w, r)
			return
		}
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}
