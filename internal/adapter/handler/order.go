package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUseCase     port.OrderUseCasePort
	CustomerUseCase  port.CustomerUseCasePort
	attendantUseCase port.AttendantUseCasePort
	productUseCase   port.ProductUseCasePort
}

func NewOrderHandler(
	orderUC port.OrderUseCasePort,
	customerUC port.CustomerUseCasePort,
	attendantUC port.AttendantUseCasePort,
	productUC port.ProductUseCasePort,
) *OrderHandler {
	return &OrderHandler{
		OrderUseCase:     orderUC,
		CustomerUseCase:  customerUC,
		attendantUseCase: attendantUC,
		productUseCase:   productUC,
	}
}

type StartOrderRequest struct {
	CustomerCPF string `json:"customerCPF"`
	AttendantID uint   `json:"attendantID"`
}

type StartOrderResponse struct {
	ID uint `json:"id"`
}

func (h *OrderHandler) StartOrder(c *gin.Context) {
	var request StartOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if _, err := h.attendantUseCase.GetAttendantById(uint(request.AttendantID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer, err := h.CustomerUseCase.GetCustomerByCPF(request.CustomerCPF)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.OrderUseCase.StartOrder(customer.ID, request.AttendantID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := StartOrderResponse{
		ID: order.ID,
	}

	c.JSON(http.StatusCreated, response)
}

type AddItemToOrderRequest struct {
	ProductID uint `json:"productID"`
	Quantity  int  `json:"quantity"`
}

type AddItemToOrderResponse struct {
	ID         uint    `json:"id"`
	Amount     float64 `json:"amount"`
	ItemsTotal int     `json:"itemsTotal"`
}

func (h *OrderHandler) AddItemToOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var request AddItemToOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := h.productUseCase.GetProductById(request.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	orderUpdated, err := h.OrderUseCase.AddItemToOrder(order, product, request.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(orderUpdated)

	c.JSON(http.StatusCreated, response)
}

type OrderResponse struct {
	ID            uint                 `json:"id"`
	CustomerID    uint                 `json:"customerID"`
	CustomerCPF   string               `json:"customerCPF"`
	CustomerName  string               `json:"customerName"`
	AttendantID   uint                 `json:"attendantID"`
	AttendantName string               `json:"attendantName"`
	Amount        string               `json:"amount"`
	ItemsTotal    int                  `json:"itemsTotal"`
	Status        string               `json:"status"`
	Payment       OrderPaymentResponse `json:"payment"`
	Items         []OrderItemResponse
}

type OrderPaymentResponse struct {
	Status string `json:"status"`
	Method string `json:"method"`
}

type OrderItemResponse struct {
	ProductID   uint    `json:"productID"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

func mapOrderResponse(order *entity.Order) OrderResponse {
	response := OrderResponse{
		ID:            order.ID,
		CustomerCPF:   order.Customer.CPF,
		CustomerName:  order.Customer.Name,
		AttendantID:   order.Attendant.ID,
		AttendantName: order.Attendant.Name,
		Amount:        fmt.Sprintf("%.2f", order.Amount()),
		ItemsTotal:    order.ItemsQuantity(),
		Status:        order.Status.ToString(),
		Payment: OrderPaymentResponse{
			Status: order.Payment.Status.ToString(),
			Method: order.Payment.Method.ToString(),
		},
		Items: []OrderItemResponse{},
	}

	for _, item := range order.Items {
		response.Items = append(response.Items, OrderItemResponse{
			ProductID:   item.Product.ID,
			ProductName: item.Product.Name,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		})
	}
	return response
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(order)

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) ConfirmationOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err = h.OrderUseCase.ConfirmationOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(order)

	c.JSON(http.StatusOK, response)
}

type PaymentOrderRequest struct {
	PaymentMethod string `json:"paymentMethod"`
}

func (h *OrderHandler) PaymentOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request PaymentOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err = h.OrderUseCase.PaymentOrder(order, request.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(order)

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) InPreparationOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err = h.OrderUseCase.InPreparationOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(order)

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) ReadyForDeliveryOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err = h.OrderUseCase.ReadyForDeliveryOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(order)

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) SentForDeliveryOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err = h.OrderUseCase.SentForDeliveryOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(order)

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) DeliveredOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err = h.OrderUseCase.DeliveredOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := mapOrderResponse(order)

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	order, err = h.OrderUseCase.CancelOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := mapOrderResponse(order)
	c.JSON(http.StatusOK, response)
}
