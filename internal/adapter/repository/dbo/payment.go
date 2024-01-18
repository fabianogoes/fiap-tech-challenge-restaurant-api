package dbo

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/core/domain"
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

func (p *Payment) ToEntity() *domain.Payment {
	return &domain.Payment{
		ID:     p.ID,
		Date:   p.Date,
		Method: p.toPaymentMethod(),
		Status: p.toPaymentStatus(),
		Value:  p.Value,
	}
}

func ToPaymentDBO(payment *domain.Payment) Payment {
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

func (p *Payment) toPaymentStatus() domain.PaymentStatus {
	switch p.Status {
	case "PENDING":
		return domain.PaymentStatusPending
	case "PAID":
		return domain.PaymentStatusPaid
	case "REVERSED":
		return domain.PaymentStatusReversed
	case "CANCELED":
		return domain.PaymentStatusCanceled
	case "PAYMENT_ERROR":
		return domain.PaymentStatusError
	default:
		return domain.PaymentStatusNone
	}
}

func (p *Payment) toPaymentMethod() domain.PaymentMethod {
	switch p.Method {
	case "CREDIT_CARD":
		return domain.PaymentMethodCreditCard
	case "DEBIT_CARD":
		return domain.PaymentMethodDebitCard
	case "MONEY":
		return domain.PaymentMethodMoney
	case "PIX":
		return domain.PaymentMethodPIX
	default:
		return domain.PaymentMethodNone
	}
}
