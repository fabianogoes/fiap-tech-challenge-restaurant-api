package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var productIDSuccess = uint(1)
var ProductSuccess = &entities.Product{
	ID:    productIDSuccess,
	Name:  "Test Product",
	Price: 100_00,
	Category: &entities.Category{
		ID:        1,
		Name:      "Test Category",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestProductService_CreateProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("CreateProduct", mock.Anything, mock.Anything, mock.Anything).Return(ProductSuccess, nil)

	service := NewProductService(repository)
	product, err := service.CreateProduct(ProductSuccess.Name, ProductSuccess.Price, ProductSuccess.Category.ID)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestProductService_GetProductById(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("GetProductById", mock.Anything).Return(ProductSuccess, nil)

	service := NewProductService(repository)
	product, err := service.GetProductById(productIDSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestProductService_GetProducts(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("GetProducts").Return([]*entities.Product{ProductSuccess}, nil)

	service := NewProductService(repository)
	products, err := service.GetProducts()
	assert.NoError(t, err)
	assert.NotNil(t, products)
}

func TestProductService_UpdateProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("UpdateProduct", mock.Anything).Return(ProductSuccess, nil)

	service := NewProductService(repository)
	product, err := service.UpdateProduct(ProductSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestProductService_DeleteProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("DeleteProduct", mock.Anything).Return(nil)

	service := NewProductService(repository)
	err := service.DeleteProduct(productIDSuccess)
	assert.NoError(t, err)
}
