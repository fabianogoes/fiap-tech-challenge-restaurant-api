package port

import "github.com/fiap/challenge-gofood/internal/core/domain"

// Primary ports to Customer

type ProductUseCasePort interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	GetProduct(id int64) (*domain.Product, error)
	GetProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int64) error
}

// Secondary ports to Product

type ProductRepositoryPort interface {
	CreateProduct(product *domain.Product) (*domain.Product, error)
	GetProduct(id int64) (*domain.Product, error)
	GetProducts() ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
	DeleteProduct(id int64) error
}
