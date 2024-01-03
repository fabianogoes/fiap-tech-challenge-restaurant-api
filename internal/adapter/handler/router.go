package handler

import "github.com/gin-gonic/gin"

type Router struct {
	*gin.Engine
}

func NewRouter(
	customerHandler *CustomerHandler,
	attendantHandler *AttendantHandler,
) (*Router, error) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the API GoFood",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "UP",
		})
	})

	customers := router.Group("/customers")
	{
		customers.GET("/", customerHandler.GetCustomers)
		customers.GET("/:id", customerHandler.GetCustomer)
		customers.POST("/", customerHandler.CreateCustomer)
		customers.PUT("/:id", customerHandler.UpdateCustomer)
		customers.DELETE("/:id", customerHandler.DeleteCustomer)
	}

	attendants := router.Group("/attendants")
	{
		attendants.GET("/", attendantHandler.GetAttendants)
		attendants.GET("/:id", attendantHandler.GetAttendant)
		attendants.POST("/", attendantHandler.CreateAttendant)
		attendants.PUT("/:id", attendantHandler.UpdateAttendant)
		attendants.DELETE("/:id", attendantHandler.DeleteAttendant)
	}

	return &Router{
		router,
	}, nil
}
