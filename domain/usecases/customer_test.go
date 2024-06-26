package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var customerIDSuccess = uint(1)
var customer = &entities.Customer{
	ID:        customerIDSuccess,
	Name:      "Test Customer",
	Email:     "test@test.com",
	CPF:       "12345678901",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestCustomerService_CreateCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("CreateCustomer", mock.Anything).Return(customer, nil)
	service := NewCustomerService(repository)

	createCustomer, err := service.CreateCustomer(customer.Name, customer.Email, customer.CPF)
	assert.NoError(t, err)
	assert.NotNil(t, createCustomer)
}

func TestCustomerService_GetCustomerById(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("GetCustomerById", mock.Anything).Return(customer, nil)

	service := NewCustomerService(repository)
	customer, err := service.GetCustomerById(customerIDSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerService_GetCustomerByCPF(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("GetCustomerByCPF", mock.Anything).Return(customer, nil)

	service := NewCustomerService(repository)

	customerResponse, err := service.GetCustomerByCPF(customer.CPF)
	assert.NoError(t, err)
	assert.NotNil(t, customerResponse)
}

func TestCustomerService_GetCustomers(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("GetCustomers").Return([]*entities.Customer{customer}, nil)

	service := NewCustomerService(repository)
	customers, err := service.GetCustomers()
	assert.NoError(t, err)
	assert.NotNil(t, customers)
}

func TestCustomerService_UpdateCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("UpdateCustomer", mock.Anything).Return(customer, nil)

	service := NewCustomerService(repository)
	updateCustomer, err := service.UpdateCustomer(customer)
	assert.NoError(t, err)
	assert.NotNil(t, updateCustomer)
}

func TestCustomerService_DeleteCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("DeleteCustomer", mock.Anything).Return(nil)

	service := NewCustomerService(repository)
	err := service.DeleteCustomer(customerIDSuccess)
	assert.NoError(t, err)
}
