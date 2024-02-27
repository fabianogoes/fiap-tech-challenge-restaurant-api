package dbo

import (
	"github.com/fabianogoes/fiap-challenge/entities"
	"gorm.io/gorm"
)

// Attendant is a Database Object for attendant
type Attendant struct {
	gorm.Model
	Name string `gorm:"unique"`
}

// ToEntity converts Attendant DBO to entities.Attendant
func (a *Attendant) ToEntity() *entities.Attendant {
	return &entities.Attendant{
		ID:        a.ID,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		Name:      a.Name,
	}
}

// ToDBO converts entities.Attendant to Attendant DBO
func ToAttendantDBO(a *entities.Attendant) *Attendant {
	return &Attendant{
		Model: gorm.Model{
			ID:        a.ID,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
		},
		Name: a.Name,
	}
}
