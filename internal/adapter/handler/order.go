package handler

import (
	"net/http"

	"github.com/fiap/challenge-gofood/internal/core/port"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUseCase     port.OrderUseCasePort
	CustomerUseCase  port.CustomerUseCasePort
	attendantUseCase port.AttendantUseCasePort
}

func NewOrderHandler(
	orderUC port.OrderUseCasePort,
	customerUC port.CustomerUseCasePort,
	attendantUC port.AttendantUseCasePort,

) *OrderHandler {
	return &OrderHandler{
		OrderUseCase:     orderUC,
		CustomerUseCase:  customerUC,
		attendantUseCase: attendantUC,
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
