package controllers

import (
	"net/http"

	"github.com/abdul-rehman-d/library-api/db"
	"github.com/gin-gonic/gin"
)

func CheckoutBook(c *gin.Context) {
	id, exists := c.GetQuery("book_id")

	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Parameter 'book_id' is required.",
		})
		return
	}

	book, err := FindBookByID(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	if book.Quantity <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "this book is unavailable.",
		})
		return
	}

	result := db.DB.Model(&book).Update("Quantity", book.Quantity-1)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.IndentedJSON(http.StatusAccepted, book)
}
