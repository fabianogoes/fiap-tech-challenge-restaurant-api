package interfaces

import "github.com/fiap/challenge-gofood/entities"

// Primary ports to Customer

type ProductUseCasePort interface {
	CreateProduct(name string, price float64, categoryID uint) (*entities.Product, error)
	GetProductById(id uint) (*entities.Product, error)
	GetProducts() ([]*entities.Product, error)
	UpdateProduct(product *entities.Product) (*entities.Product, error)
	DeleteProduct(id uint) error
}

// Secondary ports to Product

type ProductRepositoryPort interface {
	CreateProduct(name string, price float64, categoryID uint) (*entities.Product, error)
	GetProductById(id uint) (*entities.Product, error)
	GetProductByName(name string) (*entities.Product, error)
	GetProducts() ([]*entities.Product, error)
	UpdateProduct(product *entities.Product) (*entities.Product, error)
	DeleteProduct(id uint) error
}
