package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("some_secret_key")

func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwtKey)
}
