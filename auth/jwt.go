package auth

import (
	"errors"
	"os"
	"time"

	"github.com/gabrielgaspar447/go-blog-api/constants"
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
			return nil, errors.New("Invalid token")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("Invalid token")
	}

	return uint(claims["userId"].(float64)), nil
}

func AuthHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if tokenLen := len(token); tokenLen < 7 || token[:7] != "Bearer " {
		c.JSON(constants.HTTP_Unauthorized, gin.H{"error": constants.MissingToken})
		c.Abort()
		return
	}

	userId, err := verifyJWT(token[7:])

	if err != nil {
		c.JSON(constants.HTTP_Unauthorized, gin.H{"error": constants.InvalidToken})
		c.Abort()
		return
	}

	c.Set("userId", userId)
	c.Next()
}
