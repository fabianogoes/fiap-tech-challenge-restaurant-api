package repository

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Date   time.Time
	Method string
	Status string
	Value  float64
}

func (p *Payment) ToModel() *domain.Payment {
	return &domain.Payment{
		ID:     p.ID,
		Date:   p.Date,
		Method: mapPaymentMethod(p.Method),
		Status: mapPaymentStatus(p.Status),
		Value:  p.Value,
	}
}

func mapPatmentEntity(payment *domain.Payment) Payment {
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

func mapPaymentMethod(method string) domain.PaymentMethod {
	switch method {
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

func mapPaymentStatus(status string) domain.PaymentStatus {
	switch status {
	case "PENDING":
		return domain.PaymentStatusPending
	case "PAID":
		return domain.PaymentStatusPaid
	case "CANCELED":
		return domain.PaymentStatusCanceled
	case "PAYMENT_ERROR":
		return domain.PaymentStatusError
	default:
		return domain.PaymentStatusNone
	}
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (p *PaymentRepository) GetPaymentById(id uint) (*domain.Payment, error) {
	var payment Payment
	if err := p.db.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}

	return payment.ToModel(), nil
}

func (p *PaymentRepository) UpdatePayment(payment *domain.Payment) (*domain.Payment, error) {
	var err error
	paymentEntity := mapPatmentEntity(payment)

	if err = p.db.Save(&paymentEntity).Error; err != nil {
		return nil, err
	}

	return paymentEntity.ToModel(), nil
}
