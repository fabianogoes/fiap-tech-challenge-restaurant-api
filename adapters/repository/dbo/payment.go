package dbo

import (
	"github.com/fabianogoes/fiap-challenge/entities"
	"time"

	"gorm.io/gorm"
)

// Payment is a Database Object for payment
type Payment struct {
	gorm.Model
	Date   time.Time
	Method string
	Status string
	Value  float64
}

func (p *Payment) ToEntity() *entities.Payment {
	return &entities.Payment{
		ID:     p.ID,
		Date:   p.Date,
		Method: p.toPaymentMethod(),
		Status: p.toPaymentStatus(),
		Value:  p.Value,
	}
}

func ToPaymentDBO(payment *entities.Payment) Payment {
	return Payment{
		Model: gorm.Model{
			ID:        payment.ID,
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
		},
		Date:   payment.Date,
		Method: payment.Method.ToString(),
		Status: payment.Status.ToString(),
		Value:  payment.Value,
	}
}

func (p *Payment) toPaymentStatus() entities.PaymentStatus {
	switch p.Status {
	case "PENDING":
		return entities.PaymentStatusPending
	case "PAID":
		return entities.PaymentStatusPaid
	case "REVERSED":
		return entities.PaymentStatusReversed
	case "CANCELED":
		return entities.PaymentStatusCanceled
	case "PAYMENT_ERROR":
		return entities.PaymentStatusError
	default:
		return entities.PaymentStatusNone
	}
}

func (p *Payment) toPaymentMethod() entities.PaymentMethod {
	switch p.Method {
	case "CREDIT_CARD":
		return entities.PaymentMethodCreditCard
	case "DEBIT_CARD":
		return entities.PaymentMethodDebitCard
	case "MONEY":
		return entities.PaymentMethodMoney
	case "PIX":
		return entities.PaymentMethodPIX
	default:
		return entities.PaymentMethodNone
	}
}
