package service

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"github.com/fiap/challenge-gofood/internal/core/port"
)

type ProductService struct {
	productRepository port.ProductRepositoryPort
}

func NewProductService(rep port.ProductRepositoryPort) *ProductService {
	return &ProductService{
		productRepository: rep,
	}
}

func (c *ProductService) CreateProduct(name string, price float64, categoryID uint) (*domain.Product, error) {
	return c.productRepository.CreateProduct(name, price, categoryID)
}

func (c *ProductService) GetProductById(id uint) (*domain.Product, error) {
	return c.productRepository.GetProductById(id)
}

func (c *ProductService) GetProducts() ([]*domain.Product, error) {
	return c.productRepository.GetProducts()
}

func (c *ProductService) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	return c.productRepository.UpdateProduct(product)
}

func (c *ProductService) DeleteProduct(id uint) error {
	return c.productRepository.DeleteProduct(id)
}
