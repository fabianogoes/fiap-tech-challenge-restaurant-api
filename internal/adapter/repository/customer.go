package repository

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Nome  string
	Email string
	CPF   string
}
