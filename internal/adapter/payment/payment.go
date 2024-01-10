package payment

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

type PaymentClientUseCase struct {
}

func NewPaymentClientUseCase() *PaymentClientUseCase {
	return &PaymentClientUseCase{}
}

func (p *PaymentClientUseCase) Pay(order *entity.Order) error {
	fmt.Printf("Order %d paid by method %s\n", order.ID, order.Payment.Method.ToString())
	return nil
}

func (p *PaymentClientUseCase) Reverse(order *entity.Order) error {
	fmt.Printf("Order %d payment reversed\n", order.ID)
	return nil
}
