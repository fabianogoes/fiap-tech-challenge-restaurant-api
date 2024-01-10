package repository

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string
	Email string
	CPF   string `gorm:"unique"`
}

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

	result, err := c.GetCustomerByCPF(customer.CPF)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CustomerRepository) GetCustomerByCPF(cpf string) (*entity.Customer, error) {
	var result Customer
	if err := c.db.Where("cpf = ?", cpf).First(&result).Error; err != nil {
		return nil, fmt.Errorf("error to find customer with cpf %s - %v", cpf, err)
	}

	return &entity.Customer{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CPF:       result.CPF,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *CustomerRepository) GetCustomerById(id uint) (*entity.Customer, error) {
	var result Customer
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, fmt.Errorf("error to find customer with id %d - %v", id, err)
	}

	return &entity.Customer{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CPF:       result.CPF,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *CustomerRepository) GetCustomers() ([]*entity.Customer, error) {
	var results []*Customer
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var customers []*entity.Customer
	for _, result := range results {
		customers = append(customers, &entity.Customer{
			ID:        result.ID,
			Name:      result.Name,
			Email:     result.Email,
			CPF:       result.CPF,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}

	return customers, nil
}

func (c *CustomerRepository) UpdateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	var result Customer
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
	if err := c.db.Delete(&Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}
