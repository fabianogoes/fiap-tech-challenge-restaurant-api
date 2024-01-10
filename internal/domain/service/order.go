package service

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
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

func (os *OrderService) StartOrder(customerID uint, attendantID uint) (*entity.Order, error) {
	return os.orderRepository.StartOrder(
		customerID,
		attendantID,
		entity.OrderStatusStarted.ToString(),
		entity.PaymentStatusPending.ToString(),
	)
}

func (os *OrderService) GetOrderById(id uint) (*entity.Order, error) {
	return os.orderRepository.GetOrderById(id)
}

func (os *OrderService) AddItemToOrder(order *entity.Order, product *entity.Product, quantity int) (*entity.Order, error) {

	order.Amount += product.Price * float64(quantity)
	order.ItemsTotal += quantity
	order.Status = entity.OrderStatusAddingItems

	order.Items = append(order.Items, &entity.OrderItem{
		Product:   product,
		Quantity:  quantity,
		UnitPrice: product.Price,
	})

	return os.orderRepository.AddItemToOrder(order, product, quantity)
}

func (os *OrderService) ConfirmationOrder(order *entity.Order) (*entity.Order, error) {
	order.Status = entity.OrderStatusConfirmed
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrder(order *entity.Order, paymentMethod string) (*entity.Order, error) {
	payment, err := os.paymentUseCase.GetPaymentById(order.Payment.ID)
	if err != nil {
		return nil, err
	}

	if err := os.paymentClient.Pay(order); err != nil {
		order.Status = entity.OrderStatusPaymentError
		payment.Status = entity.PaymentStatusError
	} else {
		order.Status = entity.OrderStatusPaid
		payment.Status = entity.PaymentStatusPaid
	}

	payment.Method = entity.ToPaymentMethod(paymentMethod)
	_, err = os.paymentUseCase.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	order.Payment = payment
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) InPreparationOrder(order *entity.Order) (*entity.Order, error) {
	order.Status = entity.OrderStatusInPreparation
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ReadyForDeliveryOrder(order *entity.Order) (*entity.Order, error) {
	order.Status = entity.OrderStatusReadyForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) SentForDeliveryOrder(order *entity.Order) (*entity.Order, error) {
	order.Status = entity.OrderStatusSentForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) DeliveredOrder(order *entity.Order) (*entity.Order, error) {
	order.Status = entity.OrderStatusDelivered
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) CancelOrder(order *entity.Order) (*entity.Order, error) {
	fmt.Printf("CancelOrder Payment Status: %s\n", order.Payment.Status.ToString())
	if order.Payment.Status == entity.PaymentStatusPaid {
		payment, _ := os.paymentUseCase.GetPaymentById(order.Payment.ID)
		if err := os.paymentClient.Reverse(order); err != nil {
			order.Status = entity.OrderStatusPaymentError
			payment.Status = entity.PaymentStatusError
		} else {
			payment.Status = entity.PaymentStatusReversed
			_, err = os.paymentUseCase.UpdatePayment(payment)
			if err != nil {
				return nil, err
			}

			order.Payment = payment
		}
	}

	order.Status = entity.OrderStatusCanceled
	return os.orderRepository.UpdateOrder(order)
}
