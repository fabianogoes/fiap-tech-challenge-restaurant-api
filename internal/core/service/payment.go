package service

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"github.com/fiap/challenge-gofood/internal/core/port"
)

type PaymentService struct {
	paymentRepository port.PaymentRepositoryPort
}

func NewPaymentService(rep port.PaymentRepositoryPort) *PaymentService {
	return &PaymentService{
		paymentRepository: rep,
	}
}

func (c *PaymentService) GetPaymentById(id uint) (*domain.Payment, error) {
	return c.paymentRepository.GetPaymentById(id)
}

func (c *PaymentService) UpdatePayment(payment *domain.Payment) (*domain.Payment, error) {
	return c.paymentRepository.UpdatePayment(payment)
}
