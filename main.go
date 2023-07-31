package main

import (
	"errors"
	"net/http"
	"strconv"

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

func findBook(id string) (*book, error) {
	for _, b := range books {
		if b.ID == id {
			return &b, nil
		}
	}

	return nil, errors.New("book not found")
}

func getBookByID(c *gin.Context) {
	id := c.Param("bookId")

	b, err := findBook(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, b)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data provided.",
		})
		return
	}

	index, err := strconv.ParseInt(books[len(books)-1].ID, 10, 0)

	if err != nil {
		index = int64(len(books))
	}

	newBook.ID = strconv.FormatInt(index+1, 10)

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	r := gin.Default()

	// routes
	r.GET("/books", getBooks)
	r.POST("/books", createBook)
	r.GET("/books/:bookId", getBookByID)

	r.Run("localhost:8000")
}
