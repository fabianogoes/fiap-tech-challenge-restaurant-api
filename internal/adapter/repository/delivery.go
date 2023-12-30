package repository

import (
	"time"

	"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	OrderID int
	Order   Order
	Date    time.Time
	Status  string
}
