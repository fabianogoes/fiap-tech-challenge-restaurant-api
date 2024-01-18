package port

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
)

// Primary ports to Customer

type CustomerUseCasePort interface {
	CreateCustomer(nome string, email string, cpf string) (*domain.Customer, error)
	GetCustomerById(id uint) (*domain.Customer, error)
	GetCustomerByCPF(cpf string) (*domain.Customer, error)
	GetCustomers() ([]*domain.Customer, error)
	UpdateCustomer(customer *domain.Customer) (*domain.Customer, error)
	DeleteCustomer(id uint) error
}

// Secondary ports to Customer

type CustomerRepositoryPort interface {
	CreateCustomer(customer *domain.Customer) (*domain.Customer, error)
	GetCustomerByCPF(cpf string) (*domain.Customer, error)
	GetCustomerById(id uint) (*domain.Customer, error)
	GetCustomers() ([]*domain.Customer, error)
	UpdateCustomer(customer *domain.Customer) (*domain.Customer, error)
	DeleteCustomer(id uint) error
}
