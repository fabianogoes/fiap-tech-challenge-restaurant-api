package repository

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (p *PaymentRepository) GetPaymentById(id uint) (*entities.Payment, error) {
	var payment dbo.Payment

	if err := p.db.Where("id = ?", id).First(&payment).Error; err != nil {
		return nil, err
	}

	return payment.ToEntity(), nil
}

func (p *PaymentRepository) UpdatePayment(payment *entities.Payment) (*entities.Payment, error) {
	var err error
	paymentEntity := dbo.ToPaymentDBO(payment)

	if err = p.db.Save(&paymentEntity).Error; err != nil {
		return nil, err
	}

	return paymentEntity.ToEntity(), nil
}
