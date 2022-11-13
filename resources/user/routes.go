package user

import (
	"github.com/gabrielgaspar447/go-blog-api/auth"
	"github.com/gin-gonic/gin"
)

func LoadUserRoutes(app *gin.RouterGroup) {
	app.POST("/signup", signupHandler)
	app.POST("/login", loginHandler)
	app.GET("/list", listUsersHandler)
	app.GET("/:id", getUserByIdHandler)
	app.DELETE("/me", auth.AuthHandler, deleteUserHandler)
}
