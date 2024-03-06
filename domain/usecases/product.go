package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
)

type ProductService struct {
	productRepository ports.ProductRepositoryPort
}

func NewProductService(rep ports.ProductRepositoryPort) *ProductService {
	return &ProductService{
		productRepository: rep,
	}
}

func (c *ProductService) CreateProduct(name string, price float64, categoryID uint) (*entities.Product, error) {
	return c.productRepository.CreateProduct(name, price, categoryID)
}

func (c *ProductService) GetProductById(id uint) (*entities.Product, error) {
	return c.productRepository.GetProductById(id)
}

func (c *ProductService) GetProducts() ([]*entities.Product, error) {
	return c.productRepository.GetProducts()
}

func (c *ProductService) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	return c.productRepository.UpdateProduct(product)
}

func (c *ProductService) DeleteProduct(id uint) error {
	return c.productRepository.DeleteProduct(id)
}
