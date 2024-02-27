package delivery

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/entities"
)

type ClientAdapter struct {
}

func NewDeliveryClientAdapter() *ClientAdapter {
	return &ClientAdapter{}
}

func (d *ClientAdapter) Deliver(order *entities.Order) error {
	fmt.Printf("Order %d delivered\n", order.ID)
	return nil
}
