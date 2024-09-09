package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
)

type CustomerService struct {
	customerRepository ports.CustomerRepositoryPort
}

func NewCustomerService(cr ports.CustomerRepositoryPort) *CustomerService {
	return &CustomerService{
		customerRepository: cr,
	}
}

func (c *CustomerService) CreateCustomer(nome string, email string, cpf string) (*entities.Customer, error) {
	customer := entities.NewCustomer(nome, email, cpf)

	return c.customerRepository.CreateCustomer(customer)
}

func (c *CustomerService) GetCustomerById(id uint) (*entities.Customer, error) {
	return c.customerRepository.GetCustomerById(id)
}

func (c *CustomerService) GetCustomerByCPF(cpf string) (*entities.Customer, error) {
	return c.customerRepository.GetCustomerByCPF(cpf)
}

func (c *CustomerService) GetCustomers() ([]*entities.Customer, error) {
	return c.customerRepository.GetCustomers()
}

func (c *CustomerService) UpdateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	return c.customerRepository.UpdateCustomer(customer)
}

func (c *CustomerService) DeleteCustomer(id uint) error {
	return c.customerRepository.DeleteCustomer(id)
}

func (c *CustomerService) InactivationByCPF(cpf string) error {
	customer, err := c.GetCustomerByCPF(cpf)
	if err != nil {
		return err
	}

	return c.DeleteCustomer(customer.ID)
}
