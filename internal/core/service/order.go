package service

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/core/domain"
	"github.com/fiap/challenge-gofood/internal/core/port"
)

const (
	NOT_POSSIBLE_WITHOUT_PAYMENT = "It is not possible to %s order: %d not paid"
	NOT_POSSIBLE_WITHOUT_ITEMS   = "It is not possible to %s the order: %d without items"
)

type OrderService struct {
	orderRepository     port.OrderRepositoryPort
	customerRepository  port.CustomerRepositoryPort
	attendantRepository port.AttendantRepositoryPort
	paymentUseCase      port.PaymentUseCasePort
	paymentClient       port.PaymentClientPort
	deliveryClient      port.DeliveryClientPort
	deliveryRepository  port.DeliveryRepositoryPort
}

func NewOrderService(
	orderRepo port.OrderRepositoryPort,
	customerRepo port.CustomerRepositoryPort,
	attendantRepo port.AttendantRepositoryPort,
	paymentUC port.PaymentUseCasePort,
	paymentClient port.PaymentClientPort,
	deliveryClient port.DeliveryClientPort,
	deliveryRepo port.DeliveryRepositoryPort,
) *OrderService {
	return &OrderService{
		orderRepository:     orderRepo,
		customerRepository:  customerRepo,
		attendantRepository: attendantRepo,
		paymentUseCase:      paymentUC,
		paymentClient:       paymentClient,
		deliveryClient:      deliveryClient,
		deliveryRepository:  deliveryRepo,
	}
}

func (os *OrderService) StartOrder(customerID uint, attendantID uint) (*domain.Order, error) {
	var err error
	customer, err := os.customerRepository.GetCustomerById(customerID)
	if err != nil {
		return nil, err
	}

	attendant, err := os.attendantRepository.GetAttendantById(attendantID)
	if err != nil {
		return nil, err
	}

	order, err := domain.NewOrder(customer, attendant)
	if err != nil {
		return nil, err
	}

	return os.orderRepository.CreateOrder(order)
}

func (os *OrderService) GetOrderById(id uint) (*domain.Order, error) {
	return os.orderRepository.GetOrderById(id)
}

func (os *OrderService) AddItemToOrder(order *domain.Order, product *domain.Product, quantity int) (*domain.Order, error) {

	order.AddItem(product, quantity)
	order.Status = domain.OrderStatusAddingItems

	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) RemoveItemFromOrder(order *domain.Order, idItem uint) (*domain.Order, error) {
	_, err := os.orderRepository.GetOrderItemById(idItem)
	if err != nil {
		return nil, fmt.Errorf("item not found with id %d - %v", idItem, err)
	}

	switch order.Status {
	case domain.OrderStatusSentForDelivery,
		domain.OrderStatusDelivered:
		return nil, fmt.Errorf("It is not possible to REMOVE ITEM the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

	if order.Payment.Status == domain.PaymentStatusPaid {
		payment, _ := os.paymentUseCase.GetPaymentById(order.Payment.ID)
		if err := os.paymentClient.Reverse(order); err != nil {
			order.Status = domain.OrderStatusPaymentError
			payment.Status = domain.PaymentStatusError
		} else {
			payment.Status = domain.PaymentStatusReversed
			_, err = os.paymentUseCase.UpdatePayment(payment)
			if err != nil {
				return nil, err
			}

			order.Payment = payment
		}
	}

	os.orderRepository.RemoveItemFromOrder(idItem)

	order.Status = domain.OrderStatusAddingItems
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ConfirmationOrder(order *domain.Order) (*domain.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "CONFIRM", order.ID)
	}

	order.Status = domain.OrderStatusConfirmed
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrder(order *domain.Order, paymentMethod string) (*domain.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "PAY", order.ID)
	}

	if order.Status != domain.OrderStatusConfirmed {
		return nil, fmt.Errorf("It is not possible to PAY the order: %d without CONFIRMED", order.ID)
	}

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

func (os *OrderService) DeliveredOrder(order *domain.Order) (*domain.Order, error) {
	if order.Payment.Status != domain.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != domain.OrderStatusSentForDelivery {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without SENT FOR DELIVERY", order.ID)
	}

	delivery, err := os.deliveryRepository.GetDeliveryById(order.Delivery.ID)
	if err != nil {
		return nil, err
	}

	if err := os.deliveryClient.Deliver(order); err != nil {
		order.Status = domain.OrderStatusDeliveryError
		delivery.Status = domain.DeliveryStatusError
	} else {
		order.Status = domain.OrderStatusDelivered
		delivery.Status = domain.DeliveryStatusDelivered
	}

	_, err = os.deliveryRepository.UpdateDelivery(delivery)
	if err != nil {
		return nil, err
	}

	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) InPreparationOrder(order *domain.Order) (*domain.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "PREPARE", order.ID)
	}

	if order.Status != domain.OrderStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "PREPARE", order.ID)
	}

	order.Status = domain.OrderStatusInPreparation
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ReadyForDeliveryOrder(order *domain.Order) (*domain.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Payment.Status != domain.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != domain.OrderStatusInPreparation {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without PREPARE", order.ID)
	}

	order.Status = domain.OrderStatusReadyForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) SentForDeliveryOrder(order *domain.Order) (*domain.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "DELIVERY", order.ID)
	}

	if order.Payment.Status != domain.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != domain.OrderStatusReadyForDelivery {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without READY FOR DELIVERY", order.ID)
	}

	delivery, err := os.deliveryRepository.GetDeliveryById(order.Delivery.ID)
	if err != nil {
		return nil, err
	}

	delivery.Status = domain.DeliveryStatusSent
	_, err = os.deliveryRepository.UpdateDelivery(delivery)
	if err != nil {
		return nil, err
	}

	order.Status = domain.OrderStatusSentForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) CancelOrder(order *domain.Order) (*domain.Order, error) {
	switch order.Status {
	case domain.OrderStatusSentForDelivery,
		domain.OrderStatusDelivered:
		return nil, fmt.Errorf("It is not possible to cancel the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

	if order.Payment.Status == domain.PaymentStatusPaid {
		payment, _ := os.paymentUseCase.GetPaymentById(order.Payment.ID)
		if err := os.paymentClient.Reverse(order); err != nil {
			order.Status = domain.OrderStatusPaymentError
			payment.Status = domain.PaymentStatusError
		} else {
			payment.Status = domain.PaymentStatusReversed
			_, err = os.paymentUseCase.UpdatePayment(payment)
			if err != nil {
				return nil, err
			}

			order.Payment = payment
		}
	}

	order.Status = domain.OrderStatusCanceled
	return os.orderRepository.UpdateOrder(order)
}
