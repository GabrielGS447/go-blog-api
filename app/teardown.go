package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gabrielgs447/go-blog-api/database"
)

func Teardown(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Stopping server...")
	server.Shutdown(ctx)
	fmt.Println("Closing database connection...")
	database.Disconnect()

}
