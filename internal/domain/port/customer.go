package port

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

// Primary ports to Customer

type CustomerUseCasePort interface {
	CreateCustomer(nome string, email string, cpf string) (*entity.Customer, error)
	GetCustomerById(id uint) (*entity.Customer, error)
	GetCustomerByCPF(cpf string) (*entity.Customer, error)
	GetCustomers() ([]*entity.Customer, error)
	UpdateCustomer(customer *entity.Customer) (*entity.Customer, error)
	DeleteCustomer(id uint) error
}

// Secondary ports to Customer

type CustomerRepositoryPort interface {
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	GetCustomerByCPF(cpf string) (*entity.Customer, error)
	GetCustomerById(id uint) (*entity.Customer, error)
	GetCustomers() ([]*entity.Customer, error)
	UpdateCustomer(customer *entity.Customer) (*entity.Customer, error)
	DeleteCustomer(id uint) error
}
