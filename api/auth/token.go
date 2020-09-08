package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenIsValid ...
func TokenIsValid(r *http.Request) error {
	bearerToken := GetTokenFromRequest(r)

	token, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		PrintToken(claims)
	}

	return nil
}

// CreateToken creates a jwt token
func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Token expires after 2 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// GetTokenFromRequest ...
func GetTokenFromRequest(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, "")[1]
	}

	return ""
}

// GetUserIDFromToken ...
func GetUserIDFromToken(r *http.Request) (int, error) {
	bearerToken := GetTokenFromRequest(r)

	token, err := jwt.Parse(bearerToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%d", claims["user_id"]), 10, 32)

		if err != nil {
			return 0, err
		}

		return int(userID), nil
	}

	return 0, nil
}

// PrintToken on terminal
func PrintToken(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
