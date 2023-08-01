package controllers

import (
	"net/http"
	"strconv"

	"github.com/abdul-rehman-d/library-api/db"
	"github.com/abdul-rehman-d/library-api/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	result := db.DB.Find(&books)

	if err := result.Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func CreateBook(c *gin.Context) {
	var bookFromJson models.Book

	if err := c.BindJSON(&bookFromJson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data provided.",
		})
		return
	}

	newBook := models.Book{
		Title:    bookFromJson.Title,
		Author:   bookFromJson.Author,
		Quantity: bookFromJson.Quantity,
	}

	db.DB.Create(&newBook)

	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("bookId")

	parsedID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID: ID must be an integer.",
		})
		return
	}

	book := models.Book{}
	result := db.DB.Find(&book, uint(parsedID))

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong.",
		})
		return
	}

	if book.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Book Not Found.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
