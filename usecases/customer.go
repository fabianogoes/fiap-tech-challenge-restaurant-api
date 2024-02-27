package usecases

import (
	"github.com/fiap/challenge-gofood/entities"
	"github.com/fiap/challenge-gofood/interfaces"
)

type CustomerService struct {
	customerRepository interfaces.CustomerRepositoryPort
}

func NewCustomerService(cr interfaces.CustomerRepositoryPort) *CustomerService {
	return &CustomerService{
		customerRepository: cr,
	}
}

func (c *CustomerService) CreateCustomer(nome string, email string, cpf string) (*entities.Customer, error) {
	customer, err := entities.NewCustomer(nome, email, cpf)
	if err != nil {
		panic(err)
	}

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
