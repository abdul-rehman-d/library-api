package main

import (
	"github.com/abdul-rehman-d/go-first-api/controllers"
	"github.com/abdul-rehman-d/go-first-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDB()

	r := gin.Default()

	// routes
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", controllers.CreateBook)
	// r.GET("/books/:bookId", getBookByID)

	// r.PATCH("/books/checkout", checkoutBook)

	r.Run("localhost:8000")
}
