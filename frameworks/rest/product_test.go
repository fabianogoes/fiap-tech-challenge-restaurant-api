package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/usecases"
	"github.com/fabianogoes/fiap-challenge/frameworks/rest/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProduct_GetProducts(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("GetProducts").Return([]*entities.Product{domain.ProductSuccess}, nil)

	setup := SetupTest()
	setup.GET("/products", handler.GetProducts)
	request, err := http.NewRequest("GET", fmt.Sprintf("/products"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_GetProductsNoContent(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("GetProducts").Return([]*entities.Product{}, nil)

	setup := SetupTest()
	setup.GET("/products", handler.GetProducts)
	request, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_GetProductsInternalServerError(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("GetProducts").Return([]*entities.Product{}, errors.New("error"))

	setup := SetupTest()
	setup.GET("/products", handler.GetProducts)
	request, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_GetProductById(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, nil)

	setup := SetupTest()
	setup.GET("/products/:id", handler.GetProductById)
	request, err := http.NewRequest("GET", "/products/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_GetProductByIdInternalServerError(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, errors.New("error"))

	setup := SetupTest()
	setup.GET("/products/:id", handler.GetProductById)
	request, err := http.NewRequest("GET", "/products/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_GetProductByIdBadRequest(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, errors.New("error"))

	setup := SetupTest()
	setup.GET("/products/:id", handler.GetProductById)
	request, err := http.NewRequest("GET", "/products/x", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_CreateProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("CreateProduct", mock.Anything, mock.Anything, mock.Anything).
		Return(domain.ProductSuccess, nil)

	payload := dto.CreateProductRequest{Name: "test", Price: 1, CategoryID: 1}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/products", handler.CreateProduct)
	request, err := http.NewRequest("POST", "/products", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_CreateProductInternalServerError(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("CreateProduct", mock.Anything, mock.Anything, mock.Anything).
		Return(domain.ProductSuccess, errors.New("error"))

	payload := dto.CreateProductRequest{Name: "test", Price: 1, CategoryID: 1}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/products", handler.CreateProduct)
	request, err := http.NewRequest("POST", "/products", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_CreateProductStatusBadRequest(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("CreateProduct", mock.Anything, mock.Anything, mock.Anything).
		Return(domain.ProductSuccess, errors.New("error"))

	setup := SetupTest()
	setup.POST("/products", handler.CreateProduct)
	request, err := http.NewRequest("POST", "/products", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_UpdateProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("UpdateProduct", mock.Anything).Return(domain.ProductSuccess, nil)
	repository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, nil)

	payload := dto.UpdateProductRequest{Name: "test", Price: 1, Category: "test"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/products/:id", handler.UpdateProduct)
	request, err := http.NewRequest("POST", "/products/1", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_UpdateProductBadRequestJson(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("UpdateProduct", mock.Anything).Return(domain.ProductSuccess, nil)
	repository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, errors.New("error"))

	setup := SetupTest()
	setup.POST("/products/:id", handler.UpdateProduct)
	request, err := http.NewRequest("POST", "/products/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_UpdateProductInternalServerError(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, errors.New("error"))

	payload := dto.UpdateProductRequest{Name: "test", Price: 1, Category: "test"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/products/:id", handler.UpdateProduct)
	request, err := http.NewRequest("POST", "/products/1", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_UpdateProductBadRequestId(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	setup := SetupTest()
	setup.POST("/products/:id", handler.UpdateProduct)
	request, err := http.NewRequest("POST", "/products/x", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_DeleteProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("DeleteProduct", mock.Anything).Return(nil)

	setup := SetupTest()
	setup.DELETE("/products/:id", handler.DeleteProduct)
	request, err := http.NewRequest("DELETE", "/products/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestProduct_DeleteProductInternalServerError(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	repository.On("DeleteProduct", mock.Anything).Return(errors.New("error"))

	setup := SetupTest()
	setup.DELETE("/products/:id", handler.DeleteProduct)
	request, err := http.NewRequest("DELETE", "/products/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	repository.AssertCalled(t, "DeleteProduct", mock.Anything)
}

func TestProduct_DeleteProductStatusBadRequest(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	useCase := usecases.NewProductService(repository)
	handler := NewProductHandler(useCase)

	setup := SetupTest()
	setup.DELETE("/products/:id", handler.DeleteProduct)
	request, err := http.NewRequest("DELETE", "/products/x", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	repository.AssertNotCalled(t, "DeleteProduct", mock.Anything)
}
