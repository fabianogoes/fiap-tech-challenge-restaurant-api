package domain

import (
	"time"
)

type ProductType int

const (
	Snack ProductType = iota
	Drink
	Combo
)

type Product struct {
	ID         int
	Nome       string
	Price      float64
	Quantity   int64
	Type       ProductType
	CategoryID int
	Category   Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Category is an entity that represents a category of product
type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
