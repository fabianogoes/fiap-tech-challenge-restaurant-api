package payment

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

type PaymentClientAdapter struct {
}

func NewPaymentClientAdapter() *PaymentClientAdapter {
	return &PaymentClientAdapter{}
}

func (p *PaymentClientAdapter) Pay(order *entity.Order) error {
	fmt.Printf("Order %d paid by method %s\n", order.ID, order.Payment.Method.ToString())
	return nil
}

func (p *PaymentClientAdapter) Reverse(order *entity.Order) error {
	fmt.Printf("Order %d payment reversed\n", order.ID)
	return nil
}
