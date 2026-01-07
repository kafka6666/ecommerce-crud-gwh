package middleware

import (
	"log"
	"net/http"
)

func ExtraLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example extra logging functionality
		log.Println("Extra Logger Middleware: Received request for", r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Extra Logger Middleware: Completed request for", r.URL.Path)
	})
}
