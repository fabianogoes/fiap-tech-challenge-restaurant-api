package repository

import (
	"github.com/fiap/challenge-gofood/internal/adapter/repository/dbo"
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (p *PaymentRepository) GetPaymentById(id uint) (*domain.Payment, error) {
	var payment dbo.Payment

	if err := p.db.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}

	return payment.ToEntity(), nil
}

func (p *PaymentRepository) UpdatePayment(payment *domain.Payment) (*domain.Payment, error) {
	var err error
	paymentEntity := dbo.ToPaymentDBO(payment)

	if err = p.db.Save(&paymentEntity).Error; err != nil {
		return nil, err
	}

	return paymentEntity.ToEntity(), nil
}
