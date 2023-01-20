package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gabrielgaspar447/go-blog-api/app"
)

func main() {
	server, err := app.Setup()
	if err != nil {
		panic(err)
	}

	go start(server)

	fmt.Println("Application started on port " + server.Addr)

	waitForShutdownSignal()

	app.Teardown(server)

	fmt.Println("Application gracefully stopped.")

	os.Exit(0)
}

func start(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func waitForShutdownSignal() {
	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
