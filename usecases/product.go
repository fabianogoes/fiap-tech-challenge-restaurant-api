package usecases

import (
	"github.com/fabianogoes/fiap-challenge/entities"
	"github.com/fabianogoes/fiap-challenge/interfaces"
)

type ProductService struct {
	productRepository interfaces.ProductRepositoryPort
}

func NewProductService(rep interfaces.ProductRepositoryPort) *ProductService {
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
