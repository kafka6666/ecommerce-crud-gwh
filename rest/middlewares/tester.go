package middlewares

import (
	"log"
	"net/http"
)

func Tester(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Tester Middleware: printed")
		next.ServeHTTP(w, r)
	})
}
