package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func checkSecretKey() error {
	if len(jwtKey) == 0 {
		return errors.New("JWT_SECRET is not set in environment variables")
	}
	return nil
}

func GenerateToken(id uint, username string, role string, expiresIn time.Duration) (string, error) {
	if err := checkSecretKey(); err != nil {
		return "", err
	}

	if expiresIn == 0 {
		expiresIn = time.Hour * 72
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   id,
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(expiresIn).Unix(),
	})

	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	if len(jwtKey) == 0 {
		return nil, errors.New("JWT_SECRET is not set in environment variables")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or invalid claims")
	}

	return claims, nil
}
