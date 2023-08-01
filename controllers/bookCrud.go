package controllers

import (
	"errors"
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

func FindBookByID(id string) (*models.Book, error) {
	parsedID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return nil, errors.New("invalid id: id must be an integer")
	}

	book := models.Book{}
	result := db.DB.Find(&book, uint(parsedID))

	if result.Error != nil {
		return nil, errors.New("something went wrong")
	}

	if book.ID == 0 {
		return nil, errors.New("book not found")
	}

	return &book, nil
}

func GetBookByID(c *gin.Context) {
	id := c.Param("bookId")

	book, err := FindBookByID(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
