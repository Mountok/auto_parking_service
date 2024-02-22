package middleware

import (
	"log"
	"net/http"
)

func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s Method: %s",r.RequestURI,r.Method)
		next.ServeHTTP(w,r)
	})
}