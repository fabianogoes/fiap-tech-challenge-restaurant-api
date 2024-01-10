package service

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
)

type CustomerService struct {
	customerRepository port.CustomerRepositoryPort
}

func NewCustomerService(cr port.CustomerRepositoryPort) *CustomerService {
	return &CustomerService{
		customerRepository: cr,
	}
}

func (c *CustomerService) CreateCustomer(nome string, email string, cpf string) (*entity.Customer, error) {
	customer, err := entity.NewCustomer(nome, email, cpf)
	if err != nil {
		panic(err)
	}

	return c.customerRepository.CreateCustomer(customer)
}

func (c *CustomerService) GetCustomerById(id uint) (*entity.Customer, error) {
	return c.customerRepository.GetCustomerById(id)
}

func (c *CustomerService) GetCustomerByCPF(cpf string) (*entity.Customer, error) {
	return c.customerRepository.GetCustomerByCPF(cpf)
}

func (c *CustomerService) GetCustomers() ([]*entity.Customer, error) {
	return c.customerRepository.GetCustomers()
}

func (c *CustomerService) UpdateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	return c.customerRepository.UpdateCustomer(customer)
}

func (c *CustomerService) DeleteCustomer(id uint) error {
	return c.customerRepository.DeleteCustomer(id)
}
