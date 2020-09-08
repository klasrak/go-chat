package middlewares

import (
	"errors"
	"net/http"

	"github.com/klasrak/go-chat/api/auth"
	"github.com/klasrak/go-chat/api/responses"
)

// JSON set content type to application/json
func JSON(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		n(w, r)
	}
}

// JWT middleware
func JWT(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenIsValid(r)

		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		n(w, r)
	}
}
