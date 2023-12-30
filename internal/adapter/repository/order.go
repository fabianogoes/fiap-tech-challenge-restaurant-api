package repository

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID    int
	Customer      Customer
	AttendantID   int
	Attendant     Attendant
	Date          time.Time
	Status        string
	PaymentStatus string
	Amount        float64
	Items         []*OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   int
	Order     Order
	ProductID int
	Product   Product
	Quantity  int64
	UnitPrice float64
}
