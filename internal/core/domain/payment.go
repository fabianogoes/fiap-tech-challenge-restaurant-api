package domain

import "time"

type PaymentStatus int

const (
	PaymentStatusPending PaymentStatus = iota
	PaymentStatusPaid
	PaymentStatusCanceled
)

func (ps PaymentStatus) ToString() string {
	return [...]string{"PENDING", "PAID", "CANCELED"}[ps]
}

func (ps PaymentStatus) ToPaymentStatus(status string) PaymentStatus {
	switch status {
	case "PENDING":
		return PaymentStatusPending
	case "PAID":
		return PaymentStatusPaid
	case "CANCELED":
		return PaymentStatusCanceled
	default:
		return PaymentStatusPending
	}
}

type PaymentMethod int

const (
	PaymentMethodCreditCard PaymentMethod = iota
	PaymentMethodDebitCard
	PaymentMethodMoney
)

func (pm PaymentMethod) ToString() string {
	return [...]string{"CREDIT_CARD", "DEBIT_CARD", "MONEY"}[pm]
}

func (pm PaymentMethod) ToPaymentMethod(method string) PaymentMethod {
	switch method {
	case "CREDIT_CARD":
		return PaymentMethodCreditCard
	case "DEBIT_CARD":
		return PaymentMethodDebitCard
	case "MONEY":
		return PaymentMethodMoney
	default:
		return PaymentMethodMoney
	}
}

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
