package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/klasrak/go-chat/api/models"
)

// Server ...
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize ...
func (s *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error
	DBUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)

	s.DB, err = gorm.Open(DBDriver, DBUrl)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", DBDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("connected to %s database\n\n", DBDriver)
	}

	s.DB.AutoMigrate(&models.Role{}, &models.User{})
	s.Router = mux.NewRouter()
	s.InitRoutes()
}

// Run ...
func (s *Server) Run(addr string) {
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}
