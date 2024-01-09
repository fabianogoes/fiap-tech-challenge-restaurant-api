package payment

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/core/domain"
)

type PaymentClientUseCase struct {
}

func NewPaymentClientUseCase() *PaymentClientUseCase {
	return &PaymentClientUseCase{}
}

func (p *PaymentClientUseCase) Pay(order *domain.Order) error {
	fmt.Printf("Order %d paid by method %s\n", order.ID, order.Payment.Method.ToString())
	return nil
}
