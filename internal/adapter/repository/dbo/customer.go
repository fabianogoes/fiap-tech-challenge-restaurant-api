package dbo

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

// Customer is a Database Object for customer
type Customer struct {
	gorm.Model
	Name  string
	Email string
	CPF   string `gorm:"unique"`
}

// ToEntity converts Customer DBO to entity.Customer
func (c *Customer) ToEntity() *entity.Customer {
	return &entity.Customer{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Name:      c.Name,
		Email:     c.Email,
		CPF:       c.CPF,
	}
}

// ToDBO converts entity.Customer to Customer DBO
func ToCustomerDBO(c *entity.Customer) *Customer {
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
