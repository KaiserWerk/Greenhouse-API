package middleware

import (
	"net/http"
	"os"

	"github.com/KaiserWerk/Greenhouse-Manager/internal/config"
)

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		envKey := os.Getenv(config.EnvKey)
		headerKey := r.Header.Get(config.HeaderKey)

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
