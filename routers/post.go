package routers

import (
	"github.com/gabrielgs447/go-blog-api/auth"
	"github.com/gabrielgs447/go-blog-api/database"
	"github.com/gabrielgs447/go-blog-api/handlers"
	"github.com/gabrielgs447/go-blog-api/services"
	"github.com/gin-gonic/gin"
)

func LoadPostRoutes(app *gin.Engine) {
	handler := makePostHandler()
	post := app.Group("/post")
	post.POST("/create", auth.AuthHandler, handler.Create)
	post.GET("/list", handler.List)
	post.GET("/search", handler.Search)
	post.GET("/:id", handler.GetById)
	post.PATCH("/:id", auth.AuthHandler, handler.Update)
	post.DELETE("/:id", auth.AuthHandler, handler.Delete)
}

func makePostHandler() handlers.PostHandlerInterface {
	repository := database.NewPostRepository()
	service := services.NewPostService(repository)
	return handlers.NewPostHandler(service)
}
