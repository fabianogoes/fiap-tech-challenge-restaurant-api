package dbo

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

// Customer is a Database Object for customer
type Customer struct {
	gorm.Model
	Name  string
	Email string
	CPF   string `gorm:"unique"`
}

// ToEntity converts Customer DBO to domain.Customer
func (c *Customer) ToEntity() *domain.Customer {
	return &domain.Customer{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Name:      c.Name,
		Email:     c.Email,
		CPF:       c.CPF,
	}
}

// ToDBO converts domain.Customer to Customer DBO
func ToCustomerDBO(c *domain.Customer) *Customer {
	return &Customer{
		Model: gorm.Model{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		Name:  c.Name,
		Email: c.Email,
		CPF:   c.CPF,
	}
}
