package auth

import (
	"os"
	"time"

	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = os.Getenv("JWT_SECRET")

// SignJWT signs a JWT token
func SignJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}
