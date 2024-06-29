package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestProductService_CreateProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("CreateProduct", mock.Anything, mock.Anything, mock.Anything).Return(domain.ProductSuccess, nil)

	service := NewProductService(repository)
	product, err := service.CreateProduct(domain.ProductSuccess.Name, domain.ProductSuccess.Price, domain.ProductSuccess.Category.ID)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestProductService_GetProductById(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, nil)

	service := NewProductService(repository)
	product, err := service.GetProductById(domain.ProductSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestProductService_GetProducts(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("GetProducts").Return([]*entities.Product{domain.ProductSuccess}, nil)

	service := NewProductService(repository)
	products, err := service.GetProducts()
	assert.NoError(t, err)
	assert.NotNil(t, products)
}

func TestProductService_UpdateProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("UpdateProduct", mock.Anything).Return(domain.ProductSuccess, nil)

	service := NewProductService(repository)
	product, err := service.UpdateProduct(domain.ProductSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestProductService_DeleteProduct(t *testing.T) {
	repository := new(domain.ProductRepositoryMock)
	repository.On("DeleteProduct", mock.Anything).Return(nil)

	service := NewProductService(repository)
	err := service.DeleteProduct(domain.ProductSuccess.ID)
	assert.NoError(t, err)
}
