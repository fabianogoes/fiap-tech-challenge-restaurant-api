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

func (os *OrderService) GetOrderById(id uint) (*domain.Order, error) {
	return os.orderRepository.GetOrderById(id)
}

func (os *OrderService) AddItemToOrder(order *domain.Order, product *domain.Product, quantity int) (*domain.Order, error) {

	order.Amount += product.Price * float64(quantity)
	order.ItemsTotal += quantity
	order.Status = domain.OrderStatusAddingItems

	order.Items = append(order.Items, &domain.OrderItem{
		Product:   product,
		Quantity:  quantity,
		UnitPrice: product.Price,
	})

	return os.orderRepository.AddItemToOrder(order, product, quantity)
}
