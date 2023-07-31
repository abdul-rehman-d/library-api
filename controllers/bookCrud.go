package controllers

import (
	"net/http"

	"github.com/abdul-rehman-d/go-first-api/db"
	"github.com/abdul-rehman-d/go-first-api/models"
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
