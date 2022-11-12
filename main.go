package main

import (
	"github.com/gabrielgaspar447/go-blog-api/config"
	"github.com/gabrielgaspar447/go-blog-api/db"
)

func main() {
	config.LoadEnvs()
	db.Connect(true)
}
