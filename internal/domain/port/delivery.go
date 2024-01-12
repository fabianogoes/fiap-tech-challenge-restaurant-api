package port

import "github.com/fiap/challenge-gofood/internal/domain/entity"

type DeliveryClientPort interface {
	Deliver(order *entity.Order) error
}

type DeliveryRepositoryPort interface {
	GetDeliveryById(id uint) (*entity.Delivery, error)
	CreateDelivery(delivery *entity.Delivery) (*entity.Delivery, error)
	UpdateDelivery(delivery *entity.Delivery) (*entity.Delivery, error)
}
