package routers

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/handlers"
	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(app *gin.Engine) {
	user := app.Group("/user")

	user.POST("/signup", handlers.SignupHandler)
	user.POST("/login", handlers.LoginHandler)
	user.GET("/list", handlers.ListUsersHandler)
	user.GET("/:id", handlers.GetUserByIdHandler)
	user.DELETE("/me", auth.AuthHandler, handlers.DeleteUserHandler)
}
