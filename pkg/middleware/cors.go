package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CORSMiddleware() func(http.Handler) http.Handler {
	ch := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"},
		MaxAge:         3600,
	})

	return func(h http.Handler) http.Handler {
		return ch.Handler(h)
	}
}
