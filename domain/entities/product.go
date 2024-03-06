package entities

import (
	"time"
)

type ProductType int

type Product struct {
	ID        uint
	Name      string
	Price     float64
	Category  *Category
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
