package rest

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/fabianogoes/fiap-challenge/frameworks/rest/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	UseCase ports.CustomerUseCasePort
	Config  *entities.Config
}

func NewCustomerHandler(useCase ports.CustomerUseCasePort, config *entities.Config) *CustomerHandler {
	return &CustomerHandler{
		UseCase: useCase,
		Config:  config,
	}
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

func (h *CustomerHandler) SignIn(c *gin.Context) {
	var err error
	var request dto.TokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	customer, err := h.UseCase.GetCustomerByCPF(request.CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	tokenResponse, err := generateToken(customer, h.Config.TokenSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": tokenResponse.AccessToken,
		"type":        tokenResponse.Type,
		"exp":         tokenResponse.ExpiresAt,
	})
}

func generateToken(customer *entities.Customer, secret string) (*dto.TokenResponse, error) {
	secretKey := []byte(secret)

	expiresAt := time.Now().Add(time.Hour * 1).Unix() // Token expira em 1 hora
	claims := jwt.MapClaims{
		"sub":   customer.Email,
		"user":  customer.Name,
		"email": customer.Email,
		"exp":   expiresAt,
		"iat":   time.Now().Unix(),
		"iss":   "Restaurant Sign-in",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Erro ao assinar o token:", err)
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken: signedToken,
		Type:        "Bearer",
		ExpiresAt:   expiresAt,
	}, nil
}
