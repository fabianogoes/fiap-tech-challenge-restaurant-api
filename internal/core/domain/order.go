package domain

import "time"

type Order struct {
	ID         uint
	Customer   *Customer
	Attendant  *Attendant
	Date       time.Time
	Status     OrderStatus
	Payment    *Payment
	Amount     float64
	ItemsTotal int
	Items      []*OrderItem
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type OrderItem struct {
	ID        uint
	Order     Order
	Product   *Product
	Quantity  int
	UnitPrice float64
}

type OrderStatus int

const (
	OrderStatusStarted OrderStatus = iota
	OrderStatusAddingItems
	OrderStatusConfirmed
	OrderStatusPaid
	OrderStatusPaymentReversed
	OrderStatusPaymentError
	OrderStatusInPreparation
	OrderStatusReadyForDelivery
	OrderStatusSentForDelivery
	OrderStatusDelivered
	OrderStatusCanceled
)

func (os OrderStatus) ToString() string {
	return [...]string{
		"STARTED",
		"ADDING_ITEMS",
		"CONFIRMED",
		"PAID",
		"PAYMENT_REVERSED",
		"PAYMENT_ERROR",
		"IN_PREPARATION",
		"READY_FOR_DELIVERY",
		"SENT_FOR_DELIVERY",
		"DELIVERED",
		"CANCELED",
	}[os]
}

func (os OrderStatus) ToOrderStatus(status string) OrderStatus {
	switch status {
	case "STARTED":
		return OrderStatusStarted
	case "CONFIRMED":
		return OrderStatusConfirmed
	case "IN_PREPARATION":
		return OrderStatusInPreparation
	case "READY_FOR_DELIVERY":
		return OrderStatusReadyForDelivery
	case "PAID":
		return OrderStatusPaid
	case "PAYMENT_REVERSED":
		return OrderStatusPaymentReversed
	case "PAYMENT_ERROR":
		return OrderStatusPaymentError
	case "SENT_FOR_DELIVERY":
		return OrderStatusSentForDelivery
	case "DELIVERED":
		return OrderStatusDelivered
	case "CANCELED":
		return OrderStatusCanceled
	default:
		return OrderStatusStarted
	}
}
