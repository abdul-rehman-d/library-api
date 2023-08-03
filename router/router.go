package router

import (
	"github.com/abdul-rehman-d/library-api/controllers"
	"github.com/abdul-rehman-d/library-api/middleware"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.GetCustomCors())

	// routes
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:bookId", controllers.GetBookByID)

	router.POST("/purge", controllers.DeleteAll)

	router.PATCH("/checkout", controllers.CheckoutBook)
	router.PATCH("/checkin", controllers.CheckinBook)

	return router
}
