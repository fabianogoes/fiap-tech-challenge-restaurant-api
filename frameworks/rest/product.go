package rest

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
	"github.com/fabianogoes/fiap-challenge/frameworks/rest/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	UseCase ports.ProductUseCasePort
}

func NewProductHandler(useCase ports.ProductUseCasePort) *ProductHandler {
	return &ProductHandler{useCase}
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

	c.JSON(http.StatusOK, dto.ToProductResponses(products))
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

	c.JSON(http.StatusOK, dto.ToProductResponse(product))
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request dto.CreateProductRequest
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

	c.JSON(http.StatusCreated, dto.ToProductResponse(product))
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var request dto.UpdateProductRequest
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

	productUpdated, err := h.UseCase.UpdateProduct(product)

	c.JSON(http.StatusAccepted, dto.ToProductResponse(productUpdated))
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
