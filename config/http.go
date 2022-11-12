package config

import (
	"os"

	"github.com/gabrielgaspar447/go-blog-api/resources/user"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	user.LoadUserRoutes(app.Group("/user"))

	port := os.Getenv("SERVER_PORT")
	app.Run(":" + port)
}
