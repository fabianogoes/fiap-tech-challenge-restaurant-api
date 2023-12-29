package port

import "github.com/fiap/challenge-gofood/internal/core/domain"

// Primary ports to Customer

type CustomerUseCasePort interface {
	CreateCustomer(product *domain.Customer) (*domain.Customer, error)
	GetCustomer(id int64) (*domain.Customer, error)
	GetCustomers() ([]*domain.Customer, error)
	UpdateCustomer(product *domain.Customer) (*domain.Customer, error)
	DeleteCustomer(id int64) error
}

// Secondary ports to Customer

type CustomerRepositoryPort interface {
	CreateCustomer(product *domain.Customer) (*domain.Customer, error)
	GetCustomer(id int64) (*domain.Customer, error)
	GetCustomers() ([]*domain.Customer, error)
	UpdateCustomer(product *domain.Customer) (*domain.Customer, error)
	DeleteCustomer(id int64) error
}
