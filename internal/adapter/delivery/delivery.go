package delivery

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/core/domain"
)

type DeliveryClientAdapter struct {
}

func NewDeliveryClientAdapter() *DeliveryClientAdapter {
	return &DeliveryClientAdapter{}
}

func (d *DeliveryClientAdapter) Deliver(order *domain.Order) error {
	fmt.Printf("Order %d delivered\n", order.ID)
	return nil
}
