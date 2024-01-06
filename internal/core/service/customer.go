package service

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"github.com/fiap/challenge-gofood/internal/core/port"
)

type CustomerService struct {
	customerRepository port.CustomerRepositoryPort
}

func NewCustomerService(cr port.CustomerRepositoryPort) *CustomerService {
	return &CustomerService{
		customerRepository: cr,
	}
}

func (c *CustomerService) CreateCustomer(nome string, email string, cpf string) (*domain.Customer, error) {
	customer, err := domain.NewCustomer(nome, email, cpf)
	if err != nil {
		panic(err)
	}

	return c.customerRepository.CreateCustomer(customer)
}

func (c *CustomerService) GetCustomerById(id uint) (*domain.Customer, error) {
	return c.customerRepository.GetCustomerById(id)
}

func (c *CustomerService) GetCustomerByCPF(cpf string) (*domain.Customer, error) {
	return c.customerRepository.GetCustomerByCPF(cpf)
}

func (c *CustomerService) GetCustomers() ([]*domain.Customer, error) {
	return c.customerRepository.GetCustomers()
}

func (c *CustomerService) UpdateCustomer(customer *domain.Customer) (*domain.Customer, error) {
	return c.customerRepository.UpdateCustomer(customer)
}

func (c *CustomerService) DeleteCustomer(id uint) error {
	return c.customerRepository.DeleteCustomer(id)
}
