package config

import (
	"os"

	"github.com/gabrielgaspar447/go-blog-api/resources/post"
	"github.com/gabrielgaspar447/go-blog-api/resources/user"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	user.LoadUserRoutes(app.Group("/user"))
	post.LoadPostRoutes(app.Group("/post"))

	port := os.Getenv("SERVER_PORT")
	app.Run(":" + port)
}
