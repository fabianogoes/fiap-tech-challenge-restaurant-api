package port

import "github.com/fiap/challenge-gofood/internal/core/domain"

// Primary ports to Order

type OrderUseCasePort interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	GetOrder(id int64) (*domain.Order, error)
	GetOrders() ([]*domain.Order, error)
	UpdateOrder(order *domain.Order) (*domain.Order, error)
	DeleteOrder(id int64) error
}

// Secondary ports to Order

type OrderRepositoryPort interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	GetOrder(id int64) (*domain.Order, error)
	GetOrders() ([]*domain.Order, error)
	UpdateOrder(order *domain.Order) (*domain.Order, error)
	DeleteOrder(id int64) error
}
