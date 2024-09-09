package dbo

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/shared"
	"gorm.io/gorm"
)

type Attendant struct {
	gorm.Model
	Name string `gorm:"unique"`
}

func (a *Attendant) ToEntity(crypto *shared.Crypto) *entities.Attendant {
	return &entities.Attendant{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Name:      shared.MaskSensitiveData(crypto.DecryptAES(a.Name)),
	}
}
