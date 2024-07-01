package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method=%s, url=%s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}
