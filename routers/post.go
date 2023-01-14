package routers

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/handlers"
	"github.com/gin-gonic/gin"
)

func LoadPostRoutes(app *gin.Engine) {
	post := app.Group("/post")

	post.POST("/create", auth.AuthHandler, handlers.CreatePostHandler)
	post.GET("/list", handlers.ListPostsHandler)
	post.GET("/search", handlers.SearchPostsHandler)
	post.GET("/:id", handlers.GetPostByIdHandler)
	post.PUT("/:id", auth.AuthHandler, handlers.UpdatePostHandler)
	post.DELETE("/:id", auth.AuthHandler, handlers.DeletePostHandler)
}
