package domain

type ProductType int

const (
	Snack ProductType = iota
	Drink
	Combo
)

type Product struct {
	ID       int64
	Nome     string
	Price    float64
	Quantity int64
	Type     ProductType
}
