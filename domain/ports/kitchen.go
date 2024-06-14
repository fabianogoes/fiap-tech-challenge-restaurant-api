package ports

import "github.com/fabianogoes/fiap-challenge/domain/entities"

type KitchenClientPort interface {
	Preparation(order *entities.Order) error
	ReadyDelivery(orderID uint) error
}
