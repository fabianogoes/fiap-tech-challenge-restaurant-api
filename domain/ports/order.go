package ports

import "github.com/fabianogoes/fiap-challenge/domain/entities"

// Primary ports to Order

type OrderUseCasePort interface {
	StartOrder(customerID uint, attendantID uint) (*entities.Order, error)
	GetOrderById(id uint) (*entities.Order, error)
	GetOrders() ([]*entities.Order, error)
	AddItemToOrder(order *entities.Order, product *entities.Product, quantity int) (*entities.Order, error)
	RemoveItemFromOrder(order *entities.Order, idItem uint) (*entities.Order, error)
	ConfirmationOrder(order *entities.Order) (*entities.Order, error)
	PaymentOrder(order *entities.Order, paymentMethod string) (*entities.Order, error)
	PaymentOrderConfirmed(order *entities.Order, paymentMethod string) (*entities.Order, error)
	PaymentOrderError(order *entities.Order, paymentMethod string, errorReason string) (*entities.Order, error)
	InPreparationOrder(order *entities.Order) (*entities.Order, error)
	ReadyForDeliveryOrder(order *entities.Order) (*entities.Order, error)
	SentForDeliveryOrder(order *entities.Order) (*entities.Order, error)
	DeliveredOrder(order *entities.Order) (*entities.Order, error)
	CancelOrder(order *entities.Order) (*entities.Order, error)
}

// Secondary ports to Order

type OrderRepositoryPort interface {
	CreateOrder(entity *entities.Order) (*entities.Order, error)
	GetOrderById(id uint) (*entities.Order, error)
	GetOrders() ([]*entities.Order, error)
	UpdateOrder(order *entities.Order) (*entities.Order, error)
	RemoveItemFromOrder(idItem uint) error
	GetOrderItemById(id uint) (*entities.OrderItem, error)
}
