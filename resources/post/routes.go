package post

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gin-gonic/gin"
)

func LoadPostRoutes(app *gin.RouterGroup) {
	app.POST("/create", auth.AuthHandler, createPostHandler)
	app.GET("/list", listPostsHandler)
	app.GET("/search", searchPostsHandler)
	app.GET("/:id", getPostByIdHandler)
	app.PUT("/:id", auth.AuthHandler, updatePostHandler)
	app.DELETE("/:id", auth.AuthHandler, deletePostHandler)
}
