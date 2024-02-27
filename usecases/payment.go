package usecases

import (
	"github.com/fabianogoes/fiap-challenge/entities"
	"github.com/fabianogoes/fiap-challenge/interfaces"
)

type PaymentService struct {
	paymentRepository interfaces.PaymentRepositoryPort
}

func NewPaymentService(rep interfaces.PaymentRepositoryPort) *PaymentService {
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
