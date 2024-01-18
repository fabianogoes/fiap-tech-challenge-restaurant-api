package port

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
)

type DeliveryClientPort interface {
	Deliver(order *domain.Order) error
}

type DeliveryRepositoryPort interface {
	GetDeliveryById(id uint) (*domain.Delivery, error)
	CreateDelivery(delivery *domain.Delivery) (*domain.Delivery, error)
	UpdateDelivery(delivery *domain.Delivery) (*domain.Delivery, error)
}
