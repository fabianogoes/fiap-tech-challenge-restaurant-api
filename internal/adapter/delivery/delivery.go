package delivery

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

type DeliveryClientAdapter struct {
}

func NewDeliveryClientAdapter() *DeliveryClientAdapter {
	return &DeliveryClientAdapter{}
}

func (d *DeliveryClientAdapter) Deliver(order *entity.Order) error {
	fmt.Printf("Order %d delivered\n", order.ID)
	return nil
}
