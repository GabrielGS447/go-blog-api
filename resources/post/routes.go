package post

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gin-gonic/gin"
)

func LoadPostRoutes(app *gin.RouterGroup) {
	app.POST("/create", auth.AuthHandler, createPostHandler)
}
