package port

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
)

// Primary ports to Customer

type ProductUseCasePort interface {
	CreateProduct(name string, price float64, categoryID uint) (*domain.Product, error)
	GetProductById(id uint) (*domain.Product, error)
	GetProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
}

// Secondary ports to Product

type ProductRepositoryPort interface {
	CreateProduct(name string, price float64, categoryID uint) (*domain.Product, error)
	GetProductById(id uint) (*domain.Product, error)
	GetProductByName(name string) (*domain.Product, error)
	GetProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id uint) error
}
