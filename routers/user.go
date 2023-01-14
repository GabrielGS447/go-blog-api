package routers

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gabrielgaspar447/go-blog-api/handlers"
	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(app *gin.Engine) {
	user := app.Group("/user")

	user.POST("/signup", handlers.UserSignup)
	user.POST("/login", handlers.UserLogin)
	user.GET("/list", handlers.UserList)
	user.GET("/:id", handlers.UserGetById)
	user.DELETE("/me", auth.AuthHandler, handlers.UserDeleteSelf)
}
