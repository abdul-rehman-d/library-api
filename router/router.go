package router

import (
	"github.com/abdul-rehman-d/go-first-api/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	router := gin.Default()

	// routes
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	// r.GET("/books/:bookId", getBookByID)

	// r.PATCH("/books/checkout", checkoutBook)

	return router
}
