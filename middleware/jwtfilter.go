package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func JwtAuthFilter(c *gin.Context) {
	// Get the authorization header
	authHeader := c.GetHeader("Authorization")

	// Check if the authorization header is missing
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header is missing"})
		c.Abort()
		return
	}

	// Extract the token from the authorization header
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		// Define the secret key
		secret := []byte(os.Getenv("secret"))

		return secret, nil
	})

	// Check for token parsing errors
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	// Check if the token is valid
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	// Add the token claims to the context for later use
	claims := token.Claims.(jwt.MapClaims)
	c.Set("email", claims["email"])

	c.Next()
}
