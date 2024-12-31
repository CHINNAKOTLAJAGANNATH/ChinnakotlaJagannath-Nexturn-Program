package middleware

import (
	"net/http"
)

func ValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Invalid content type", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
