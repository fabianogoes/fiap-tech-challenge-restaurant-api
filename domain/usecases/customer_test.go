package usecases

import (
	"errors"
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCustomerService_CreateCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("CreateCustomer", mock.Anything).Return(domain.CustomerSuccess, nil)
	service := NewCustomerService(repository)

	createCustomer, err := service.CreateCustomer(domain.CustomerSuccess.Name, domain.CustomerSuccess.Email, domain.CustomerSuccess.CPF)
	assert.NoError(t, err)
	assert.NotNil(t, createCustomer)
}

func TestCustomerService_CreateCustomerError(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("CreateCustomer", mock.Anything).Return(domain.CustomerSuccess, errors.New("error"))
	service := NewCustomerService(repository)

	_, err := service.CreateCustomer(domain.CustomerSuccess.Name, domain.CustomerSuccess.Email, domain.CustomerSuccess.CPF)
	assert.Error(t, err)
}

func TestCustomerService_GetCustomerById(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	service := NewCustomerService(repository)
	customer, err := service.GetCustomerById(domain.CustomerSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, customer)
}

func TestCustomerService_GetCustomerByCPF(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("GetCustomerByCPF", mock.Anything).Return(domain.CustomerSuccess, nil)

	service := NewCustomerService(repository)

	customerResponse, err := service.GetCustomerByCPF(domain.CustomerSuccess.CPF)
	assert.NoError(t, err)
	assert.NotNil(t, customerResponse)
}

func TestCustomerService_GetCustomers(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("GetCustomers").Return([]*entities.Customer{domain.CustomerSuccess}, nil)

	service := NewCustomerService(repository)
	customers, err := service.GetCustomers()
	assert.NoError(t, err)
	assert.NotNil(t, customers)
}

func TestCustomerService_UpdateCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("UpdateCustomer", mock.Anything).Return(domain.CustomerSuccess, nil)

	service := NewCustomerService(repository)
	updateCustomer, err := service.UpdateCustomer(domain.CustomerSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, updateCustomer)
}

func TestCustomerService_DeleteCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	repository.On("DeleteCustomer", mock.Anything).Return(nil)

	service := NewCustomerService(repository)
	err := service.DeleteCustomer(domain.CustomerSuccess.ID)
	assert.NoError(t, err)
}
