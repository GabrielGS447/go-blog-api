package main

import (
	"os"

	"github.com/gabrielgaspar447/go-blog-api/app"
)

func main() {
	server, err := app.Setup()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
