package repository

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID int
	Order   Order
	Date    time.Time
	Method  string
	Status  string
	Value   float64
}
