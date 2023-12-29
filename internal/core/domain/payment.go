package domain

import "time"

type PaymentStatus int
type PaymentMethod int

const (
	PaymentStatusPending PaymentStatus = iota
	PaymentStatusPaid
	PaymentStatusCanceled
)

const (
	PaymentMethodCreditCard PaymentMethod = iota
	PaymentMethodDebitCard
	PaymentMethodMoney
)

type Payment struct {
	ID        int64
	Order     Order
	Date      time.Time
	Method    PaymentMethod
	Status    PaymentStatus
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
