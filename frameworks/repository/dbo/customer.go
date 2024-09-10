package dbo

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/shared"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string
	Email string
	CPF   string `gorm:"unique"`
}

func (c *Customer) ToEntity(crypto *shared.Crypto) *entities.Customer {
	return &entities.Customer{
		ID:        c.ID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Name:      shared.MaskSensitiveData(crypto.DecryptAES(c.Name)),
		Email:     shared.MaskSensitiveData(crypto.DecryptAES(c.Email)),
		CPF:       c.CPF,n
	}
}
