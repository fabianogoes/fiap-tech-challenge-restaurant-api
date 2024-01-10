package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	UseCase port.ProductUseCasePort
}

func NewProductHandler(useCase port.ProductUseCasePort) *ProductHandler {
	return &ProductHandler{useCase}
}

type FindProductResponse struct {
	ID       uint              `json:"id"`
	Name     string            `json:"name"`
	Price    float64           `json:"price"`
	Category *CategoryResponse `json:"category"`
}

func ProductToResponse(product *entity.Product) *FindProductResponse {
	return &FindProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Category: CategoryToResponse(product.Category),
	}
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CategoryToResponse(category *entity.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.UseCase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if len(products) == 0 {
		c.JSON(http.StatusNoContent, gin.H{
			"message": "No products found",
		})
	}

	var response []FindProductResponse
	for _, product := range products {
		response = append(response, *ProductToResponse(product))
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetProductById(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	product, err := h.UseCase.GetProductById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := *ProductToResponse(product)

	c.JSON(http.StatusOK, response)
}

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"categoryID"`
}

type CreateProductResponse struct {
	ID uint `json:"id"`
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request CreateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	fmt.Println(request)
	product, err := h.UseCase.CreateProduct(request.Name, request.Price, request.CategoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := CreateProductResponse{
		ID: product.ID,
	}

	c.JSON(http.StatusCreated, response)
}

type UpdateProductRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var request UpdateProductRequest
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

	product, err := h.UseCase.GetProductById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	product.Name = request.Name
	product.Price = request.Price
	product.Category.Name = request.Category

	_, err = h.UseCase.UpdateProduct(product)

	response := fmt.Sprintf("Product[%d] - %s updated", product.ID, request.Name)

	c.JSON(http.StatusAccepted, gin.H{
		"message": response,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	var err error
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = h.UseCase.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	response := fmt.Sprintf("Product %d deleted", id)
	println(response)

	c.JSON(http.StatusNoContent, gin.H{
		"message": response,
	})
}
