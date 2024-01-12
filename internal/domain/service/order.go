package service

import (
	"fmt"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
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
	os.orderRepository.RemoveItemFromOrder(idItem)

	order.Status = entity.OrderStatusAddingItems
	return os.orderRepository.UpdateOrder(order)
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

func (os *OrderService) DeliveredOrder(order *entity.Order) (*entity.Order, error) {
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
	order.Status = entity.OrderStatusInPreparation
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ReadyForDeliveryOrder(order *entity.Order) (*entity.Order, error) {
	order.Status = entity.OrderStatusReadyForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) SentForDeliveryOrder(order *entity.Order) (*entity.Order, error) {
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
