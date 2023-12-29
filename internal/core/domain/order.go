package domain

import "time"

type OrderStatus int

const (
	OrderStatusStarted OrderStatus = iota
	OrderStatusConfirmed
	OrderStatusInPreparation
	OrderStatusPaid
	OrderStatusDelivered
	OrderStatusCanceled
)

type Order struct {
	ID            int64
	Customer      Customer
	Attendant     Attendant
	Date          time.Time
	Status        OrderStatus
	PaymentStatus PaymentStatus
	Amount        float64
	Items         []*OrderItem
}

type OrderItem struct {
	ID        int64
	Order     Order
	Product   Product
	Quantity  int64
	UnitPrice float64
}
