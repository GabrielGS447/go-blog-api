package app

import (
	"errors"
	"os"

	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Setup() (*gin.Engine, error) {
	if err := loadENV(); err != nil {
		return nil, err
	}

	if err := connectToDatabase(); err != nil {
		return nil, err
	}

	app, err := prepareServer()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func loadENV() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func connectToDatabase() error {
	dbURL := os.Getenv("MYSQL_URL")
	if dbURL == "" {
		return errors.New("DB_URL not set")
	}

	if resetDB := os.Getenv("RESET_DB"); resetDB == "true" {
		return database.Connect(dbURL, true)
	} else {
		return database.Connect(dbURL, false)
	}
}

func prepareServer() (*gin.Engine, error) {
	app := gin.Default()
	routers.LoadUserRoutes(app)
	routers.LoadPostRoutes(app)

	return app, nil
}
