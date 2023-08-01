package router

import (
	"github.com/abdul-rehman-d/library-api/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	router := gin.Default()

	// routes
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:bookId", controllers.GetBookByID)

	// r.PATCH("/books/checkout", checkoutBook)

	return router
}
