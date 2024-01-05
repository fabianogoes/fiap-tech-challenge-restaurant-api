package repository

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
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

func (c *CustomerRepository) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
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

func (c *CustomerRepository) GetCustomerByCPF(cpf string) (*domain.Customer, error) {
	var result Customer
	if err := c.db.Where("cpf = ?", cpf).First(&result).Error; err != nil {
		return nil, err
	}

	return &domain.Customer{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CPF:       result.CPF,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *CustomerRepository) GetCustomerById(id uint) (*domain.Customer, error) {
	var result Customer
	if err := c.db.First(&result, id).Error; err != nil {
		return nil, err
	}

	return &domain.Customer{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CPF:       result.CPF,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (c *CustomerRepository) GetCustomers() ([]*domain.Customer, error) {
	var results []*Customer
	if err := c.db.Find(&results).Error; err != nil {
		return nil, err
	}

	var customers []*domain.Customer
	for _, result := range results {
		customers = append(customers, &domain.Customer{
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

func (c *CustomerRepository) UpdateCustomer(customer *domain.Customer) (*domain.Customer, error) {
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
