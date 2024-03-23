package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
// customerHandler *CustomerHandler,
// attendantHandler *AttendantHandler,
// productHandler *ProductHandler,
// orderHandler *OrderHandler,
) (*Router, error) {
	router := gin.Default()

	router.GET("/", Welcome)
	router.GET("/health", Health)
	router.GET("/env", Environment)

	customers := router.Group("/customers")
	{
		customers.GET("/", mockCustomers)
		customers.GET("/:id", mockCustomer)
		customers.POST("/", mockCustomer)
		// 	customers.GET("/", customerHandler.GetCustomers)
		// 	customers.GET("/:id", customerHandler.GetCustomer)
		// 	customers.GET("/cpf/:cpf", customerHandler.GetCustomerByCPF)
		// 	customers.POST("/", customerHandler.CreateCustomer)
		// 	customers.PUT("/:id", customerHandler.UpdateCustomer)
		// 	customers.DELETE("/:id", customerHandler.DeleteCustomer)
	}

	// attendants := router.Group("/attendants")
	// {
	// 	attendants.GET("/", attendantHandler.GetAttendants)
	// 	attendants.GET("/:id", attendantHandler.GetAttendant)
	// 	attendants.POST("/", attendantHandler.CreateAttendant)
	// 	attendants.PUT("/:id", attendantHandler.UpdateAttendant)
	// 	attendants.DELETE("/:id", attendantHandler.DeleteAttendant)
	// }

	// products := router.Group("/products")
	// {
	// 	products.GET("/", productHandler.GetProducts)
	// 	products.GET("/:id", productHandler.GetProductById)
	// 	products.POST("/", productHandler.CreateProduct)
	// 	products.PUT("/:id", productHandler.UpdateProduct)
	// 	products.DELETE("/:id", productHandler.DeleteProduct)
	// }

	// orders := router.Group("/orders")
	// {
	// 	orders.POST("/", orderHandler.StartOrder)
	// 	orders.POST("/:id/item", orderHandler.AddItemToOrder)
	// 	orders.DELETE("/:id/item/:idItem", orderHandler.RemoveItemFromOrder)
	// 	orders.GET("/:id", orderHandler.GetOrderById)
	// 	orders.GET("/", orderHandler.GetOrders)
	// 	orders.PUT("/:id/confirmation", orderHandler.ConfirmationOrder)
	// 	orders.PUT("/:id/payment", orderHandler.PaymentOrder)
	// 	orders.PUT("/:id/payment/webhook", orderHandler.PaymentWebhook)
	// 	orders.PUT("/:id/in-preparation", orderHandler.InPreparationOrder)
	// 	orders.PUT("/:id/ready-for-delivery", orderHandler.ReadyForDeliveryOrder)
	// 	orders.PUT("/:id/sent-for-delivery", orderHandler.SentForDeliveryOrder)
	// 	orders.PUT("/:id/delivered", orderHandler.DeliveredOrder)
	// 	orders.PUT("/:id/cancel", orderHandler.CancelOrder)
	// }

	return &Router{
		router,
	}, nil
}

func mockCustomers(c *gin.Context) {
	customers := []string{
		"Customer 1",
		"Customer 2",
		"Customer 3",
		"Customer 4",
	}
	c.JSON(http.StatusOK, customers)
}

func mockCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, "Customer 1")
}
