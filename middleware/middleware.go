package middleware

import (
	"log"
	"net/http"
)

// WithLogging is a logging middleware for http requests
func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Printf(
				"%s [%s] %s %s %s",
				r.RemoteAddr,
				r.Method,
				r.Host,
				r.URL.Path,
				r.URL.RawQuery,
			)
		}()
		next.ServeHTTP(w, r)
	}
}

// w.Header().Set("Content-Type", "application/json")

// WithJSONHeader adds a Content-Type application/header to the response
func withJSONHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

// WithMiddleware adds all middleware
func WithMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return withLogging(
		withJSONHeader(
			func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			},
		),
	)
}
