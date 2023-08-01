package main

import (
	"github.com/abdul-rehman-d/library-api/db"
	"github.com/abdul-rehman-d/library-api/router"
)

func main() {
	db.InitializeDB()

	router := router.Initialize()

	router.Run("localhost:8000")
}
