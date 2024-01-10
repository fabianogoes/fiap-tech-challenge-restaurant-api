package service

import (
	"github.com/fiap/challenge-gofood/internal/core/domain"
	"github.com/fiap/challenge-gofood/internal/core/port"
)

type OrderService struct {
	orderRepository    port.OrderRepositoryPort
	customerRepository port.CustomerRepositoryPort
	paymentUseCase     port.PaymentUseCasePort
	paymentClient      port.PaymentClientPort
}

func NewOrderService(
	rep port.OrderRepositoryPort,
	cr port.CustomerRepositoryPort,
	puc port.PaymentUseCasePort,
	pc port.PaymentClientPort,
) *OrderService {
	return &OrderService{
		orderRepository:    rep,
		customerRepository: cr,
		paymentUseCase:     puc,
		paymentClient:      pc,
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

func (os *OrderService) ConfirmationOrder(order *domain.Order) (*domain.Order, error) {
	order.Status = domain.OrderStatusConfirmed
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrder(order *domain.Order, paymentMethod string) (*domain.Order, error) {
	payment, err := os.paymentUseCase.GetPaymentById(order.Payment.ID)
	if err != nil {
		return nil, err
	}

	if err := os.paymentClient.Pay(order); err != nil {
		order.Status = domain.OrderStatusPaymentError
		payment.Status = domain.PaymentStatusError
	} else {
		order.Status = domain.OrderStatusPaid
		payment.Status = domain.PaymentStatusPaid
	}

	payment.Method = domain.ToPaymentMethod(paymentMethod)
	_, err = os.paymentUseCase.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	order.Payment = payment
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) InPreparationOrder(order *domain.Order) (*domain.Order, error) {
	order.Status = domain.OrderStatusInPreparation
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ReadyForDeliveryOrder(order *domain.Order) (*domain.Order, error) {
	order.Status = domain.OrderStatusReadyForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) SentForDeliveryOrder(order *domain.Order) (*domain.Order, error) {
	order.Status = domain.OrderStatusSentForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) DeliveredOrder(order *domain.Order) (*domain.Order, error) {
	order.Status = domain.OrderStatusDelivered
	return os.orderRepository.UpdateOrder(order)
}
