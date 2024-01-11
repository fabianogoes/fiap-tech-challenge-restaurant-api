package repository

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/adapter/repository/dbo"
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db,
	}
}

func (c *CustomerRepository) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	var err error

	if err = c.db.Create(customer).Error; err != nil {
		return nil, err
	}

	return c.GetCustomerByCPF(customer.CPF)
}

func (c *CustomerRepository) GetCustomerByCPF(cpf string) (*entity.Customer, error) {
	var result dbo.Customer

	if err := c.db.Where("cpf = ?", cpf).First(&result).Error; err != nil {
		return nil, fmt.Errorf("error to find customer with cpf %s - %v", cpf, err)
	}

	return result.ToEntity(), nil
}

func (c *CustomerRepository) GetCustomerById(id uint) (*entity.Customer, error) {
	var result dbo.Customer
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find customer with id %d - %v", id, err)
	}

	return result.ToEntity(), nil
}

func (c *CustomerRepository) GetCustomers() ([]*entity.Customer, error) {
	var results []*dbo.Customer
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var customers []*entity.Customer
	for _, result := range results {
		customers = append(customers, result.ToEntity())
	}

	return customers, nil
}

func (c *CustomerRepository) UpdateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	var result dbo.Customer
	if err := c.db.First(&result, customer.ID).Error; err != nil {
		return nil, err
	}

	result.Name = customer.Name
	result.Email = customer.Email
	result.CPF = customer.CPF

	if err := c.db.Save(&result).Error; err != nil {
		return nil, err
	}

	return c.GetCustomerById(customer.ID)
}

func (c *CustomerRepository) DeleteCustomer(id uint) error {
	if err := c.db.Delete(&dbo.Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}
