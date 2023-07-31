package main

import (
	"github.com/abdul-rehman-d/go-first-api/db"
	"github.com/abdul-rehman-d/go-first-api/router"
)

func main() {
	db.InitializeDB()

	router := router.Initialize()

	router.Run("localhost:8000")
}
