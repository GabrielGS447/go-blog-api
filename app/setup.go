package app

import (
	"errors"
	"net/http"
	"os"

	"github.com/gabrielgaspar447/go-blog-api/database"
	"github.com/gabrielgaspar447/go-blog-api/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Setup() (*http.Server, error) {
	if err := loadENV(); err != nil {
		return nil, err
	}

	if err := connectToDatabase(); err != nil {
		return nil, err
	}

	server, err := createServer()
	if err != nil {
		return nil, err
	}

	return server, nil
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

func createServer() (*http.Server, error) {
	router := gin.Default()
	routers.LoadUserRoutes(router)
	routers.LoadPostRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT not set")
	}

	// This wrapper allow us to gracefully shutdown the server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return server, nil
}
