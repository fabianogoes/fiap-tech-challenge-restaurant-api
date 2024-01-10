package domain

import "time"

type PaymentStatus int

const (
	PaymentStatusPending PaymentStatus = iota
	PaymentStatusPaid
	PaymentStatusCanceled
	PaymentStatusError
	PaymentStatusNone
)

func (ps PaymentStatus) ToString() string {
	return [...]string{"PENDING", "PAID", "CANCELED", "PAYMENT_ERROR", "NONE"}[ps]
}

func (ps PaymentStatus) ToPaymentStatus(status string) PaymentStatus {
	switch status {
	case "PENDING":
		return PaymentStatusPending
	case "PAID":
		return PaymentStatusPaid
	case "CANCELED":
		return PaymentStatusCanceled
	case "PAYMENT_ERROR":
		return PaymentStatusError
	default:
		return PaymentStatusNone
	}
}

type PaymentMethod int

const (
	PaymentMethodCreditCard PaymentMethod = iota
	PaymentMethodDebitCard
	PaymentMethodMoney
	PaymentMethodPIX
	PaymentMethodNone
)

func (pm PaymentMethod) ToString() string {
	return [...]string{"CREDIT_CARD", "DEBIT_CARD", "MONEY", "PIX", "NONE"}[pm]
}

func ToPaymentMethod(method string) PaymentMethod {
	switch method {
	case "CREDIT_CARD":
		return PaymentMethodCreditCard
	case "DEBIT_CARD":
		return PaymentMethodDebitCard
	case "MONEY":
		return PaymentMethodMoney
	case "PIX":
		return PaymentMethodPIX
	default:
		return PaymentMethodNone
	}
}

type Payment struct {
	ID        uint
	Order     Order
	Date      time.Time
	Method    PaymentMethod
	Status    PaymentStatus
	Value     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
