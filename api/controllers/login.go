package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/klasrak/go-chat/api/auth"
	"github.com/klasrak/go-chat/api/models"
	"github.com/klasrak/go-chat/api/responses"
	"golang.org/x/crypto/bcrypt"
)

// Login ...
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := s.SignInAndCreateToken(user.Username, user.Password)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}

// SignInAndCreateToken ...
func (s *Server) SignInAndCreateToken(username, password string) (string, error) {
	var err error

	user := models.User{}

	err = s.DB.Model(models.User{}).Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = models.CheckPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(int(user.ID))
}
