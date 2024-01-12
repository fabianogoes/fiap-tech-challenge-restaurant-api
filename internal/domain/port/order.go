package port

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

// Primary ports to Order

type OrderUseCasePort interface {
	StartOrder(customerID uint, attendantID uint) (*entity.Order, error)
	GetOrderById(id uint) (*entity.Order, error)
	AddItemToOrder(order *entity.Order, product *entity.Product, quantity int) (*entity.Order, error)
	RemoveItemFromOrder(order *entity.Order, idItem uint) (*entity.Order, error)
	ConfirmationOrder(order *entity.Order) (*entity.Order, error)
	PaymentOrder(order *entity.Order, paymentMethod string) (*entity.Order, error)
	InPreparationOrder(order *entity.Order) (*entity.Order, error)
	ReadyForDeliveryOrder(order *entity.Order) (*entity.Order, error)
	SentForDeliveryOrder(order *entity.Order) (*entity.Order, error)
	DeliveredOrder(order *entity.Order) (*entity.Order, error)
	CancelOrder(order *entity.Order) (*entity.Order, error)
}

// Secondary ports to Order

type OrderRepositoryPort interface {
	CreateOrder(entity *entity.Order) (*entity.Order, error)
	GetOrderById(id uint) (*entity.Order, error)
	UpdateOrder(order *entity.Order) (*entity.Order, error)
	RemoveItemFromOrder(idItem uint) error
}
