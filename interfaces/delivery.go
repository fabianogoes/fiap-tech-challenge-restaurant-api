package interfaces

import "github.com/fiap/challenge-gofood/entities"

type DeliveryClientPort interface {
	Deliver(order *entities.Order) error
}

type DeliveryRepositoryPort interface {
	GetDeliveryById(id uint) (*entities.Delivery, error)
	CreateDelivery(delivery *entities.Delivery) (*entities.Delivery, error)
	UpdateDelivery(delivery *entities.Delivery) (*entities.Delivery, error)
}
