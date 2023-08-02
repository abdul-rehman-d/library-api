package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/abdul-rehman-d/library-api/db"
	"github.com/abdul-rehman-d/library-api/models"
	"github.com/abdul-rehman-d/library-api/utils"
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

	if err := utils.ValidateNewBook(bookFromJson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
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

func DeleteAll(c *gin.Context) {
	var password struct {
		Password string `json:"password"`
	}

	if err := c.BindJSON(&password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Must have 'password' key in payload",
		})
		return
	}

	if password.Password != os.Getenv("LIBRARY_API_PURGE_PASSWORD") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password",
		})
		return
	}

	result := db.DB.Delete(&models.Book{}, "Title Like ?", "%%")

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d records deleted", result.RowsAffected),
	})
}
