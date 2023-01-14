package routers

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/handlers"
	"github.com/gin-gonic/gin"
)

func LoadPostRoutes(app *gin.Engine) {
	post := app.Group("/post")

	post.POST("/create", auth.AuthHandler, handlers.PostCreate)
	post.GET("/list", handlers.PostList)
	post.GET("/search", handlers.PostSearch)
	post.GET("/:id", handlers.PostGetById)
	post.PUT("/:id", auth.AuthHandler, handlers.PostUpdate)
	post.DELETE("/:id", auth.AuthHandler, handlers.PostDelete)
}
