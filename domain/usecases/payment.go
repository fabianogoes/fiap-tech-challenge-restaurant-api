package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
)

type PaymentService struct {
	paymentRepository ports.PaymentRepositoryPort
}

func NewPaymentService(rep ports.PaymentRepositoryPort) *PaymentService {
	return &PaymentService{
		paymentRepository: rep,
	}
}

func (c *PaymentService) GetPaymentById(id uint) (*entities.Payment, error) {
	return c.paymentRepository.GetPaymentById(id)
}

func (c *PaymentService) UpdatePayment(payment *entities.Payment) (*entities.Payment, error) {
	return c.paymentRepository.UpdatePayment(payment)
}
