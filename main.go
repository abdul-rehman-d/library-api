package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Golang pointers", Author: "Mr. Golang", Quantity: 10},
	{ID: "2", Title: "Goroutines", Author: "Mr. Goroutine", Quantity: 20},
	{ID: "3", Title: "Golang routers", Author: "Mr. Router", Quantity: 30},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	r := gin.Default()

	// routes
	r.GET("/books", getBooks)

	r.Run("localhost:8000")
}
