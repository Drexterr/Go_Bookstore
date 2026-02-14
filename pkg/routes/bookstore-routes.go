package routes

import (
	"github.com/Bharat/go-bookstore/middleware"
	"github.com/Bharat/go-bookstore/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterBookstoreRoutes = func(router *gin.Engine) {

	router.GET("/bookstore/", controllers.GetBook)
	router.GET("/bookstore/:Store_id", controllers.GetBookByStoreId)
	router.GET("/bookstore/:Store_id/:id", controllers.GetBookByID)
	router.GET("/bookstore/:Store_id/:id/Price", controllers.GetPriceByID)

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	auth := router.Group("/")
	auth.Use(middleware.RequireAuth)
	{
		auth.GET("/validate", controllers.Validate)

		admin := auth.Group("/")
		admin.Use(middleware.RequireRole)
		{
			admin.PUT("/bookstore/:Store_id/:id", controllers.UpdateBook)
			admin.POST("/bookstore/:Store_id/", controllers.CreateBook)
			admin.DELETE("/bookstore/:Store_id/:id", controllers.DeleteBook)
			admin.POST("/bookstore/duplicate/from/:sourceid/to/:newid", controllers.DuplicateStore)
		}

	}
}
