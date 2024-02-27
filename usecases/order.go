package usecases

import (
	"fmt"
	"github.com/fiap/challenge-gofood/entities"
	port2 "github.com/fiap/challenge-gofood/interfaces"
)

const (
	NOT_POSSIBLE_WITHOUT_PAYMENT = "It is not possible to %s order: %d not paid"
	NOT_POSSIBLE_WITHOUT_ITEMS   = "It is not possible to %s the order: %d without items"
)

type OrderService struct {
	orderRepository     port2.OrderRepositoryPort
	customerRepository  port2.CustomerRepositoryPort
	attendantRepository port2.AttendantRepositoryPort
	paymentUseCase      port2.PaymentUseCasePort
	paymentClient       port2.PaymentClientPort
	deliveryClient      port2.DeliveryClientPort
	deliveryRepository  port2.DeliveryRepositoryPort
}

func NewOrderService(
	orderRepo port2.OrderRepositoryPort,
	customerRepo port2.CustomerRepositoryPort,
	attendantRepo port2.AttendantRepositoryPort,
	paymentUC port2.PaymentUseCasePort,
	paymentClient port2.PaymentClientPort,
	deliveryClient port2.DeliveryClientPort,
	deliveryRepo port2.DeliveryRepositoryPort,
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

func (os *OrderService) StartOrder(customerID uint, attendantID uint) (*entities.Order, error) {
	var err error
	customer, err := os.customerRepository.GetCustomerById(customerID)
	if err != nil {
		return nil, err
	}

	attendant, err := os.attendantRepository.GetAttendantById(attendantID)
	if err != nil {
		return nil, err
	}

	order, err := entities.NewOrder(customer, attendant)
	if err != nil {
		return nil, err
	}

	return os.orderRepository.CreateOrder(order)
}

func (os *OrderService) GetOrderById(id uint) (*entities.Order, error) {
	return os.orderRepository.GetOrderById(id)
}

func (os *OrderService) AddItemToOrder(order *entities.Order, product *entities.Product, quantity int) (*entities.Order, error) {

	order.AddItem(product, quantity)
	order.Status = entities.OrderStatusAddingItems

	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) RemoveItemFromOrder(order *entities.Order, idItem uint) (*entities.Order, error) {
	_, err := os.orderRepository.GetOrderItemById(idItem)
	if err != nil {
		return nil, fmt.Errorf("item not found with id %d - %v", idItem, err)
	}

	switch order.Status {
	case entities.OrderStatusSentForDelivery,
		entities.OrderStatusDelivered:
		return nil, fmt.Errorf("It is not possible to REMOVE ITEM the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

	if order.Payment.Status == entities.PaymentStatusPaid {
		payment, _ := os.paymentUseCase.GetPaymentById(order.Payment.ID)
		if err := os.paymentClient.Reverse(order); err != nil {
			order.Status = entities.OrderStatusPaymentError
			payment.Status = entities.PaymentStatusError
		} else {
			payment.Status = entities.PaymentStatusReversed
			_, err = os.paymentUseCase.UpdatePayment(payment)
			if err != nil {
				return nil, err
			}

			order.Payment = payment
		}
	}

	err = os.orderRepository.RemoveItemFromOrder(idItem)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusAddingItems
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ConfirmationOrder(order *entities.Order) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "CONFIRM", order.ID)
	}

	order.Status = entities.OrderStatusConfirmed
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrder(order *entities.Order, paymentMethod string) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "PAY", order.ID)
	}

	if order.Status != entities.OrderStatusConfirmed {
		return nil, fmt.Errorf("It is not possible to PAY the order: %d without CONFIRMED", order.ID)
	}

	payment, err := os.paymentUseCase.GetPaymentById(order.Payment.ID)
	if err != nil {
		return nil, err
	}

	if err := os.paymentClient.Pay(order); err != nil {
		order.Status = entities.OrderStatusPaymentError
		payment.Status = entities.PaymentStatusError
	} else {
		order.Status = entities.OrderStatusPaid
		payment.Status = entities.PaymentStatusPaid
	}

	payment.Method = entities.ToPaymentMethod(paymentMethod)
	_, err = os.paymentUseCase.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	order.Payment = payment
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) DeliveredOrder(order *entities.Order) (*entities.Order, error) {
	if order.Payment.Status != entities.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != entities.OrderStatusSentForDelivery {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without SENT FOR DELIVERY", order.ID)
	}

	delivery, err := os.deliveryRepository.GetDeliveryById(order.Delivery.ID)
	if err != nil {
		return nil, err
	}

	if err := os.deliveryClient.Deliver(order); err != nil {
		order.Status = entities.OrderStatusDeliveryError
		delivery.Status = entities.DeliveryStatusError
	} else {
		order.Status = entities.OrderStatusDelivered
		delivery.Status = entities.DeliveryStatusDelivered
	}

	_, err = os.deliveryRepository.UpdateDelivery(delivery)
	if err != nil {
		return nil, err
	}

	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) InPreparationOrder(order *entities.Order) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "PREPARE", order.ID)
	}

	if order.Status != entities.OrderStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "PREPARE", order.ID)
	}

	order.Status = entities.OrderStatusInPreparation
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ReadyForDeliveryOrder(order *entities.Order) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Payment.Status != entities.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != entities.OrderStatusInPreparation {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without PREPARE", order.ID)
	}

	order.Status = entities.OrderStatusReadyForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) SentForDeliveryOrder(order *entities.Order) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_ITEMS, "DELIVERY", order.ID)
	}

	if order.Payment.Status != entities.PaymentStatusPaid {
		return nil, fmt.Errorf(NOT_POSSIBLE_WITHOUT_PAYMENT, "DELIVERY", order.ID)
	}

	if order.Status != entities.OrderStatusReadyForDelivery {
		return nil, fmt.Errorf("It is not possible to DELIVERY the order: %d without READY FOR DELIVERY", order.ID)
	}

	delivery, err := os.deliveryRepository.GetDeliveryById(order.Delivery.ID)
	if err != nil {
		return nil, err
	}

	delivery.Status = entities.DeliveryStatusSent
	_, err = os.deliveryRepository.UpdateDelivery(delivery)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusSentForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) CancelOrder(order *entities.Order) (*entities.Order, error) {
	switch order.Status {
	case entities.OrderStatusSentForDelivery,
		entities.OrderStatusDelivered:
		return nil, fmt.Errorf("It is not possible to cancel the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

	if order.Payment.Status == entities.PaymentStatusPaid {
		payment, _ := os.paymentUseCase.GetPaymentById(order.Payment.ID)
		if err := os.paymentClient.Reverse(order); err != nil {
			order.Status = entities.OrderStatusPaymentError
			payment.Status = entities.PaymentStatusError
		} else {
			payment.Status = entities.PaymentStatusReversed
			_, err = os.paymentUseCase.UpdatePayment(payment)
			if err != nil {
				return nil, err
			}

			order.Payment = payment
		}
	}

	order.Status = entities.OrderStatusCanceled
	return os.orderRepository.UpdateOrder(order)
}
