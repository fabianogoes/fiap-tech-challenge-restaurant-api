package port

import "github.com/fiap/challenge-gofood/internal/core/domain"

// Primary ports to Order

type OrderUseCasePort interface {
	StartOrder(customerID uint, attendantID uint) (*domain.Order, error)
	// GetOrder(id int64) (*domain.Order, error)
	// GetOrders() ([]*domain.Order, error)
	// UpdateOrder(order *domain.Order) (*domain.Order, error)
	// DeleteOrder(id int64) error
}

// Secondary ports to Order

type OrderRepositoryPort interface {
	StartOrder(customerID uint, attendantID uint, orderStatus string, paymentStatus string) (*domain.Order, error)
	// GetOrder(id int64) (*domain.Order, error)
	// GetOrders() ([]*domain.Order, error)
	// UpdateOrder(order *domain.Order) (*domain.Order, error)
	// DeleteOrder(id int64) error
}
