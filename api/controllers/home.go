package controllers

import (
	"net/http"

	"github.com/klasrak/go-chat/api/responses"
)

// Home just a nice hello world
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Hello, World!")
}
