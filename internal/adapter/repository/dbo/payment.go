package dbo

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
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

func (p *Payment) ToEntity() *entity.Payment {
	return &entity.Payment{
		ID:     p.ID,
		Date:   p.Date,
		Method: p.toPaymentMethod(),
		Status: p.toPaymentStatus(),
		Value:  p.Value,
	}
}

func ToPaymentDBO(payment *entity.Payment) Payment {
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

func (p *Payment) toPaymentStatus() entity.PaymentStatus {
	switch p.Status {
	case "PENDING":
		return entity.PaymentStatusPending
	case "PAID":
		return entity.PaymentStatusPaid
	case "REVERSED":
		return entity.PaymentStatusReversed
	case "CANCELED":
		return entity.PaymentStatusCanceled
	case "PAYMENT_ERROR":
		return entity.PaymentStatusError
	default:
		return entity.PaymentStatusNone
	}
}

func (p *Payment) toPaymentMethod() entity.PaymentMethod {
	switch p.Method {
	case "CREDIT_CARD":
		return entity.PaymentMethodCreditCard
	case "DEBIT_CARD":
		return entity.PaymentMethodDebitCard
	case "MONEY":
		return entity.PaymentMethodMoney
	case "PIX":
		return entity.PaymentMethodPIX
	default:
		return entity.PaymentMethodNone
	}
}
