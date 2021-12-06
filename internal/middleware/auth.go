package middleware

import (
	"net/http"
	"os"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		envKey := os.Getenv("GREENHOUSE_KEY")
		headerKey := r.Header.Get("X-Greenhouse-Key")

		if envKey == "" || headerKey == "" {
			http.Error(w, "empty or missing key", http.StatusUnauthorized)
			return
		}

		if envKey != headerKey {
			http.Error(w, "key is not correct", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}
