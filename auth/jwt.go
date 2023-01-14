package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gabrielgaspar447/go-blog-api/errs"
	"github.com/gabrielgaspar447/go-blog-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func SignJWT(user *models.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"userId": user.ID,
		"exp":    expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func verifyJWT(tokenString string) (uint, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrInvalidToken
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errs.ErrInvalidToken
	}

	return uint(claims["userId"].(float64)), nil
}

func AuthHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if tokenLen := len(token); tokenLen < 7 || token[:7] != "Bearer " {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrMissingToken})
		c.Abort()
		return
	}

	userId, err := verifyJWT(token[7:])

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errs.ErrInvalidToken})
		c.Abort()
		return
	}

	c.Set("userId", userId)
	c.Next()
}
