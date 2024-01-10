package port

import "github.com/fiap/challenge-gofood/internal/domain/entity"

type DeliveryClientPort interface {
	Deliver(order *entity.Order) error
}
