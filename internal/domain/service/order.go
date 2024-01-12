package service

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
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

func (os *OrderService) StartOrder(customerID uint, attendantID uint) (*entity.Order, error) {
	var err error
	customer, err := os.customerRepository.GetCustomerById(customerID)
	if err != nil {
		return nil, err
	}

	attendant, err := os.attendantRepository.GetAttendantById(attendantID)
	if err != nil {
		return nil, err
	}

	order, err := entity.NewOrder(customer, attendant)
	if err != nil {
		return nil, err
	}

	return os.orderRepository.CreateOrder(order)
}

func (os *OrderService) GetOrderById(id uint) (*entity.Order, error) {
	return os.orderRepository.GetOrderById(id)
}

func (os *OrderService) AddItemToOrder(order *entity.Order, product *entity.Product, quantity int) (*entity.Order, error) {

	order.AddItem(product, quantity)
	order.Status = entity.OrderStatusAddingItems

	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) RemoveItemFromOrder(order *entity.Order, idItem uint) (*entity.Order, error) {
	_, err := os.orderRepository.GetOrderItemById(idItem)
	if err != nil {
		return nil, fmt.Errorf("item not found with id %d - %v", idItem, err)
	}

	switch order.Status {
	case entity.OrderStatusSentForDelivery,
		entity.OrderStatusDelivered:
		return nil, fmt.Errorf("It is not possible to REMOVE ITEM the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

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

	os.orderRepository.RemoveItemFromOrder(idItem)

	order.Status = entity.OrderStatusAddingItems
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ConfirmationOrder(order *entity.Order) (*entity.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "CONFIRM", order.ID)
	}

	order.Status = entity.OrderStatusConfirmed
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrder(order *entity.Order, paymentMethod string) (*entity.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "PAY", order.ID)
	}

	if order.Status != entity.OrderStatusConfirmed {
		return nil, fmt.Errorf("It is not possible to PAY the order: %d without CONFIRMED", order.ID)
	}

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

func (os *OrderService) DeliveredOrder(order *entity.Order) (*entity.Order, error) {
	if order.Payment.Status != entity.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != entity.OrderStatusSentForDelivery {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without SENT FOR DELIVERY", order.ID)
	}

	delivery, err := os.deliveryRepository.GetDeliveryById(order.Delivery.ID)
	if err != nil {
		return nil, err
	}

	if err := os.deliveryClient.Deliver(order); err != nil {
		order.Status = entity.OrderStatusDeliveryError
		delivery.Status = entity.DeliveryStatusError
	} else {
		order.Status = entity.OrderStatusDelivered
		delivery.Status = entity.DeliveryStatusDelivered
	}

	_, err = os.deliveryRepository.UpdateDelivery(delivery)
	if err != nil {
		return nil, err
	}

	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) InPreparationOrder(order *entity.Order) (*entity.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "PREPARE", order.ID)
	}

	if order.Status != entity.OrderStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "PREPARE", order.ID)
	}

	order.Status = entity.OrderStatusInPreparation
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ReadyForDeliveryOrder(order *entity.Order) (*entity.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Payment.Status != entity.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != entity.OrderStatusInPreparation {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without PREPARE", order.ID)
	}

	order.Status = entity.OrderStatusReadyForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) SentForDeliveryOrder(order *entity.Order) (*entity.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "DELIVERY", order.ID)
	}

	if order.Payment.Status != entity.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != entity.OrderStatusReadyForDelivery {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without READY FOR DELIVERY", order.ID)
	}

	delivery, err := os.deliveryRepository.GetDeliveryById(order.Delivery.ID)
	if err != nil {
		return nil, err
	}

	delivery.Status = entity.DeliveryStatusSent
	_, err = os.deliveryRepository.UpdateDelivery(delivery)
	if err != nil {
		return nil, err
	}

	order.Status = entity.OrderStatusSentForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) CancelOrder(order *entity.Order) (*entity.Order, error) {
	switch order.Status {
	case entity.OrderStatusSentForDelivery,
		entity.OrderStatusDelivered:
		return nil, fmt.Errorf("It is not possible to cancel the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

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
