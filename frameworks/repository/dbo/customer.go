package dbo

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"gorm.io/gorm"
)

// Customer is a Database Object for customer
type Customer struct {
	gorm.Model
	Name  string
	Email string
	CPF   string `gorm:"unique"`
}

// ToEntity converts Customer DBO to entities.Customer
func (c *Customer) ToEntity() *entities.Customer {
	return &entities.Customer{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Name:      c.Name,
		Email:     c.Email,
		CPF:       c.CPF,
	}
}

// ToDBO converts entities.Customer to Customer DBO
func ToCustomerDBO(c *entities.Customer) *Customer {
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
