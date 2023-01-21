package routers

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/handlers"
	"github.com/gabrielgaspar447/go-blog-api/services"
	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(app *gin.Engine) {
	handler := makeUserHandler()
	user := app.Group("/user")
	user.POST("/signup", handler.Signup)
	user.POST("/login", handler.Login)
	user.GET("/list", handler.List)
	user.GET("/:id", handler.GetById)
	user.DELETE("/me", auth.AuthHandler, handler.DeleteSelf)
}

func makeUserHandler() handlers.UserHandlerInterface {
	repository := database.NewUserRepository()
	service := services.NewUserService(repository)
	return handlers.NewUserHandler(service)
}
