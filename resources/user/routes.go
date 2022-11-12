package user

import "github.com/gin-gonic/gin"

func LoadUserRoutes(app *gin.RouterGroup) {
	app.POST("/signup", signupHandler)
}
