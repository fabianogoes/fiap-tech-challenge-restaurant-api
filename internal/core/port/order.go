package port

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
)

// Primary ports to Order

type OrderUseCasePort interface {
	StartOrder(customerID uint, attendantID uint) (*domain.Order, error)
	GetOrderById(id uint) (*domain.Order, error)
	AddItemToOrder(order *domain.Order, product *domain.Product, quantity int) (*domain.Order, error)
	RemoveItemFromOrder(order *domain.Order, idItem uint) (*domain.Order, error)
	ConfirmationOrder(order *domain.Order) (*domain.Order, error)
	PaymentOrder(order *domain.Order, paymentMethod string) (*domain.Order, error)
	InPreparationOrder(order *domain.Order) (*domain.Order, error)
	ReadyForDeliveryOrder(order *domain.Order) (*domain.Order, error)
	SentForDeliveryOrder(order *domain.Order) (*domain.Order, error)
	DeliveredOrder(order *domain.Order) (*domain.Order, error)
	CancelOrder(order *domain.Order) (*domain.Order, error)
}

// Secondary ports to Order

type OrderRepositoryPort interface {
	CreateOrder(entity *domain.Order) (*domain.Order, error)
	GetOrderById(id uint) (*domain.Order, error)
	UpdateOrder(order *domain.Order) (*domain.Order, error)
	RemoveItemFromOrder(idItem uint) error
	GetOrderItemById(id uint) (*domain.OrderItem, error)
}
