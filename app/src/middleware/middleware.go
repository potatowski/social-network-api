package middleware

import (
	"fmt"
	"log"
	"net/http"
	"social-api/src/response"
	"social-api/src/security"
)

// Logger print the request information in console
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n Request: %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Auth is a middleware to check if the user is authenticated
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware")
		if err := security.ValidateToken(r); err != nil {
			response.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
