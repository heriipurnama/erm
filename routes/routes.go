package routes

import (
	"dbo/erm/controllers"
	"dbo/erm/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	// Protected routes
	protected := router.Group("/").Use(middleware.AuthMiddleware())

	// Customer routes
	protected.GET("/customers", controllers.GetCustomers)
	protected.GET("/customers/:id", controllers.GetCustomer)
	protected.POST("/customers", controllers.CreateCustomer)
	protected.PUT("/customers/:id", controllers.UpdateCustomer)
	protected.DELETE("/customers/:id", controllers.DeleteCustomer)
	protected.GET("/customers/search", controllers.SearchCustomer)

	// Order routes
	protected.GET("/orders", controllers.GetOrders)
	protected.GET("/orders/:id", controllers.GetOrder)
	protected.POST("/orders", controllers.CreateOrder)
	protected.PUT("/orders/:id", controllers.UpdateOrder)
	protected.DELETE("/orders/:id", controllers.DeleteOrder)
	protected.GET("/orders/search", controllers.SearchOrder)

	return router
}
