package payment

import (
	"fmt"
	"github.com/fiap/challenge-gofood/entities"
)

type ClientAdapter struct {
}

func NewPaymentClientAdapter() *ClientAdapter {
	return &ClientAdapter{}
}

func (p *ClientAdapter) Pay(order *entities.Order) error {
	fmt.Printf("Order %d paid by method %s\n", order.ID, order.Payment.Method.ToString())
	return nil
}

func (p *ClientAdapter) Reverse(order *entities.Order) error {
	fmt.Printf("Order %d payment reversed\n", order.ID)
	return nil
}
