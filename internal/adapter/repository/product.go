package repository

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Nome       string
	Price      float64
	Quantity   int64
	Type       string
	CategoryID int
	Category   Category
}

type Category struct {
	gorm.Model
	Name string
}
