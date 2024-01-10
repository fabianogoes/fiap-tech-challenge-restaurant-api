package repository

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	Date   time.Time
	Method string
	Status string
	Value  float64
}

func (p *Payment) ToModel() *entity.Payment {
	return &entity.Payment{
		ID:     p.ID,
		Date:   p.Date,
		Method: mapPaymentMethod(p.Method),
		Status: mapPaymentStatus(p.Status),
		Value:  p.Value,
	}
}

func mapPatmentEntity(payment *entity.Payment) Payment {
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

func mapPaymentMethod(method string) entity.PaymentMethod {
	switch method {
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

func mapPaymentStatus(status string) entity.PaymentStatus {
	switch status {
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

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (p *PaymentRepository) GetPaymentById(id uint) (*entity.Payment, error) {
	var payment Payment
	if err := p.db.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}

	return payment.ToModel(), nil
}

func (p *PaymentRepository) UpdatePayment(payment *entity.Payment) (*entity.Payment, error) {
	var err error
	paymentEntity := mapPatmentEntity(payment)

	if err = p.db.Save(&paymentEntity).Error; err != nil {
		return nil, err
	}

	return paymentEntity.ToModel(), nil
}
