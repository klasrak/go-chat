package controllers

import "github.com/klasrak/go-chat/api/middlewares"

// InitRoutes ...
func (s *Server) InitRoutes() {

	// Home router
	s.Router.HandleFunc("/", middlewares.JSON(s.Home)).Methods("GET")

	// Login route
	s.Router.HandleFunc("/login", middlewares.JSON(s.Login)).Methods("POST")

	// User routes
	s.Router.HandleFunc("/users", middlewares.JSON(middlewares.JWT(s.AddUser))).Methods("POST")
}
