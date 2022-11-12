package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	app := gin.Default()

	port := os.Getenv("SERVER_PORT")
	app.Run(":" + port)
}
