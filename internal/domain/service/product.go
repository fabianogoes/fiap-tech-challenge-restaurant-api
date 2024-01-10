package service

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
)

type ProductService struct {
	productRepository port.ProductRepositoryPort
}

func NewProductService(rep port.ProductRepositoryPort) *ProductService {
	return &ProductService{
		productRepository: rep,
	}
}

func (c *ProductService) CreateProduct(name string, price float64, categoryID int) (*entity.Product, error) {
	return c.productRepository.CreateProduct(name, price, categoryID)
}

func (c *ProductService) GetProductById(id uint) (*entity.Product, error) {
	return c.productRepository.GetProductById(id)
}

func (c *ProductService) GetProducts() ([]*entity.Product, error) {
	return c.productRepository.GetProducts()
}

func (c *ProductService) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	return c.productRepository.UpdateProduct(product)
}

func (c *ProductService) DeleteProduct(id uint) error {
	return c.productRepository.DeleteProduct(id)
}
