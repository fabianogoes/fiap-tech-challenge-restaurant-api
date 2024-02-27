package handler

import (
	"fmt"
	port2 "github.com/fiap/challenge-gofood/interfaces"
	"net/http"
	"strconv"

	"github.com/fiap/challenge-gofood/adapters/handler/dto"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUseCase     port2.OrderUseCasePort
	CustomerUseCase  port2.CustomerUseCasePort
	attendantUseCase port2.AttendantUseCasePort
	productUseCase   port2.ProductUseCasePort
}

func NewOrderHandler(
	orderUC port2.OrderUseCasePort,
	customerUC port2.CustomerUseCasePort,
	attendantUC port2.AttendantUseCasePort,
	productUC port2.ProductUseCasePort,
) *OrderHandler {
	return &OrderHandler{
		OrderUseCase:     orderUC,
		CustomerUseCase:  customerUC,
		attendantUseCase: attendantUC,
		productUseCase:   productUC,
	}
}

func (h *OrderHandler) StartOrder(c *gin.Context) {
	var request dto.StartOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	c.JSON(http.StatusCreated, dto.ToStartOrderResponse(order))
}

func (h *OrderHandler) AddItemToOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request dto.AddItemToOrderRequest
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

	c.JSON(http.StatusCreated, dto.ToOrderResponse(orderUpdated))
}

func (h *OrderHandler) RemoveItemFromOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		message := fmt.Errorf("error to convert order id to int - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": message.Error(),
		})
		return
	}

	itemID, err := strconv.Atoi(c.Param("iditem"))
	if err != nil {
		message := fmt.Errorf("error to convert item id to int - %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": message.Error(),
		})
		return
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	orderUpdated, err := h.OrderUseCase.RemoveItemFromOrder(order, uint(itemID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ToOrderResponse(orderUpdated))
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *OrderHandler) ConfirmationOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *OrderHandler) PaymentOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := h.OrderUseCase.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var request dto.PaymentOrderRequest
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

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *OrderHandler) InPreparationOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *OrderHandler) ReadyForDeliveryOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *OrderHandler) SentForDeliveryOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *OrderHandler) DeliveredOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	var err error
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

	c.JSON(http.StatusOK, dto.ToOrderResponse(order))
}
