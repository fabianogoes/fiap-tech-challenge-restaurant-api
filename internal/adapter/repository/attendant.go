package repository

import "gorm.io/gorm"

type Attendant struct {
	gorm.Model
	ID   int64
	Nome string
}
