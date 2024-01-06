package service

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"github.com/fiap/challenge-gofood/internal/core/port"
)

type OrderService struct {
	orderRepository    port.OrderRepositoryPort
	customerRepository port.CustomerRepositoryPort
}

func NewOrderService(rep port.OrderRepositoryPort, cr port.CustomerRepositoryPort) *OrderService {
	return &OrderService{
		orderRepository:    rep,
		customerRepository: cr,
	}
}

func (os *OrderService) StartOrder(customerID uint, attendantID uint) (*domain.Order, error) {
	return os.orderRepository.StartOrder(
		customerID,
		attendantID,
		domain.OrderStatusStarted.ToString(),
		domain.PaymentStatusPending.ToString(),
	)
}
