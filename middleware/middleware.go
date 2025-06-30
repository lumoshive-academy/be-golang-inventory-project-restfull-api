package middleware

import (
	"go-25-27/utils"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token == "" {
			utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorization")
			return
		}

		// check the token in the storage

		next.ServeHTTP(w, r)
	})
}

func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("role")
		// check level by token
		if token != "admin" {
			utils.ResponseBadRequest(w, http.StatusForbidden, "You do not have permission to access this feature.")
			return
		}

		next.ServeHTTP(w, r)
	})
}
