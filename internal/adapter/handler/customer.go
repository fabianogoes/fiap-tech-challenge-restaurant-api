package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fiap/challenge-gofood/internal/core/port"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	UseCase port.CustomerUseCasePort
}

func NewCustomerHandler(useCase port.CustomerUseCasePort) *CustomerHandler {
	return &CustomerHandler{useCase}
}

type FindCustomerResponse struct {
	ID    uint   `json:"id"`
	Nome  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

func (h *CustomerHandler) GetCustomers(c *gin.Context) {
	customers, err := h.UseCase.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if len(customers) == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No customers found",
		})
	}

	var response []FindCustomerResponse
	for _, customer := range customers {
		response = append(response, FindCustomerResponse{
			ID:    customer.ID,
			Nome:  customer.Name,
			Email: customer.Email,
			CPF:   customer.CPF,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	customer, err := h.UseCase.GetCustomerById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := FindCustomerResponse{
		ID:    customer.ID,
		Nome:  customer.Name,
		Email: customer.Email,
		CPF:   customer.CPF,
	}

	c.JSON(http.StatusOK, response)
}

type CreateCustomerRequest struct {
	Nome  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

type CreateCustomerResponse struct {
	ID uint `json:"id"`
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var request CreateCustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	customer, err := h.UseCase.CreateCustomer(request.Nome, request.Email, request.CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := CreateCustomerResponse{
		ID: customer.ID,
	}

	c.JSON(http.StatusCreated, response)
}

type UpdateCustomerRequest struct {
	Nome  string `json:"name"`
	Email string `json:"email"`
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	var request UpdateCustomerRequest
	var err error

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	customer, err := h.UseCase.GetCustomerById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	customer.Name = request.Nome
	customer.Email = request.Email

	_, err = h.UseCase.UpdateCustomer(customer)

	response := fmt.Sprintf("Customer[%d] - %s updated", customer.ID, request.Nome)

	c.JSON(http.StatusAccepted, gin.H{
		"message": response,
	})
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = h.UseCase.DeleteCustomer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := fmt.Sprintf("Customer %d deleted", id)
	println(response)

	c.JSON(http.StatusNoContent, gin.H{
		"message": response,
	})
}
