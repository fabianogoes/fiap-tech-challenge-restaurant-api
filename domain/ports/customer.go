package ports

import "github.com/fabianogoes/fiap-challenge/domain/entities"

// Primary ports to Customer

type CustomerUseCasePort interface {
	CreateCustomer(nome string, email string, cpf string) (*entities.Customer, error)
	GetCustomerById(id uint) (*entities.Customer, error)
	GetCustomerByCPF(cpf string) (*entities.Customer, error)
	GetCustomers() ([]*entities.Customer, error)
	UpdateCustomer(customer *entities.Customer) (*entities.Customer, error)
	DeleteCustomer(id uint) error
	InactivationByCPF(cpf string) error
}

// Secondary ports to Customer

type CustomerRepositoryPort interface {
	CreateCustomer(customer *entities.Customer) (*entities.Customer, error)
	GetCustomerByCPF(cpf string) (*entities.Customer, error)
	GetCustomerById(id uint) (*entities.Customer, error)
	GetCustomers() ([]*entities.Customer, error)
	UpdateCustomer(customer *entities.Customer) (*entities.Customer, error)
	DeleteCustomer(id uint) error
}
