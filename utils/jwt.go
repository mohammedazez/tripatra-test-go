package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userEmail string) (string, error) {
	secret := os.Getenv("jwt_key")
	if secret == "" {
		return "", errors.New("JWT secret not set")
	}

	claims := &jwt.StandardClaims{
		Subject:   userEmail,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
