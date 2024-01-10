package service

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
)

type PaymentService struct {
	paymentRepository port.PaymentRepositoryPort
}

func NewPaymentService(rep port.PaymentRepositoryPort) *PaymentService {
	return &PaymentService{
		paymentRepository: rep,
	}
}

func (c *PaymentService) GetPaymentById(id uint) (*entity.Payment, error) {
	return c.paymentRepository.GetPaymentById(id)
}

func (c *PaymentService) UpdatePayment(payment *entity.Payment) (*entity.Payment, error) {
	return c.paymentRepository.UpdatePayment(payment)
}
