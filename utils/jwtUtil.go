package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateJWTToken(email string) (string, error) {

	// Define the expiration time for the token
	expirationTime := time.Now().Add(time.Hour * 24) // Token expiration time 24 hours

	// Create the claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
