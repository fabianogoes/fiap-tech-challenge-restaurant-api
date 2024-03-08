package rest

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/fabianogoes/fiap-challenge/frameworks/rest/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	UseCase ports.CustomerUseCasePort
}

func NewCustomerHandler(useCase ports.CustomerUseCasePort) *CustomerHandler {
	return &CustomerHandler{useCase}
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

	c.JSON(http.StatusOK, dto.ToCustomerResponses(customers))
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

	c.JSON(http.StatusOK, dto.ToCustomerResponse(customer))
}

func (h *CustomerHandler) GetCustomerByCPF(c *gin.Context) {
	var err error
	cpf := c.Param("cpf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	customer, err := h.UseCase.GetCustomerByCPF(cpf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := dto.GetCustomerResponse{
		ID:    customer.ID,
		Nome:  customer.Name,
		Email: customer.Email,
		CPF:   customer.CPF,
	}

	c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var request dto.CreateCustomerRequest
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

	c.JSON(http.StatusCreated, dto.ToCustomerResponse(customer))
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	var request dto.UpdateCustomerRequest
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
