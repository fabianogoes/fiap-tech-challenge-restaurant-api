package usecases

import (
	"fmt"
	"log/slog"

	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/ports"
)

const (
	NotPossibleWithoutPayment = "it is not possible to %s order: %d not paid"
	NotPossibleWithoutItems   = "it is not possible to %s the order: %d without items"
)

type OrderService struct {
	orderRepository     ports.OrderRepositoryPort
	customerRepository  ports.CustomerRepositoryPort
	attendantRepository ports.AttendantRepositoryPort
	paymentUseCase      ports.PaymentUseCasePort
	deliveryClient      ports.DeliveryClientPort
	deliveryRepository  ports.DeliveryRepositoryPort
	kitchenPublisher    ports.KitchenPublisherPort
	paymentPublisher    ports.PaymentPublisherPort
}

func NewOrderService(
	orderRepo ports.OrderRepositoryPort,
	customerRepo ports.CustomerRepositoryPort,
	attendantRepo ports.AttendantRepositoryPort,
	paymentUC ports.PaymentUseCasePort,
	deliveryClient ports.DeliveryClientPort,
	deliveryRepo ports.DeliveryRepositoryPort,
	kitchenPublisher ports.KitchenPublisherPort,
	paymentPublisher ports.PaymentPublisherPort,
) *OrderService {
	return &OrderService{
		orderRepository:     orderRepo,
		customerRepository:  customerRepo,
		attendantRepository: attendantRepo,
		paymentUseCase:      paymentUC,
		deliveryClient:      deliveryClient,
		deliveryRepository:  deliveryRepo,
		kitchenPublisher:    kitchenPublisher,
		paymentPublisher:    paymentPublisher,
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

	order := entities.NewOrder(customer, attendant)

	return os.orderRepository.CreateOrder(order)
}

func (os *OrderService) GetOrderById(id uint) (*entities.Order, error) {
	return os.orderRepository.GetOrderById(id)
}

func (os *OrderService) GetOrders() ([]*entities.Order, error) {
	return os.orderRepository.GetOrders()
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

	if order.Status == entities.OrderStatusSentForDelivery || order.Status == entities.OrderStatusDelivered {
		return nil, fmt.Errorf("it is not possible to REMOVE ITEM the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

	// TODO
	//if order.Payment.Status == entities.PaymentStatusPaid {
	//	payment, _ := os.paymentUseCase.GetPaymentById(order.Payment.ID)
	//	if err := os.paymentClient.Reverse(order); err != nil {
	//		order.Status = entities.OrderStatusPaymentError
	//		payment.Status = entities.PaymentStatusError
	//	} else {
	//		payment.Status = entities.PaymentStatusReversed
	//		_, err = os.paymentUseCase.UpdatePayment(payment)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		order.Payment = payment
	//	}
	//}

	err = os.orderRepository.RemoveItemFromOrder(idItem)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusAddingItems
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ConfirmationOrder(order *entities.Order) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NotPossibleWithoutItems, "CONFIRM", order.ID)
	}

	order.Status = entities.OrderStatusConfirmed
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrder(order *entities.Order, paymentMethod string) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NotPossibleWithoutItems, "PAY", order.ID)
	}

	if (order.Status != entities.OrderStatusConfirmed) && (order.Status != entities.OrderStatusPaymentError) {
		return nil, fmt.Errorf("it is not possible to PAY the order: %d without CONFIRMED", order.ID)
	}

	if err := os.paymentPublisher.PublishPayment(order, paymentMethod); err != nil {
		slog.Error(err.Error())
		order.Status = entities.OrderStatusPaymentError
	} else {
		order.Status = entities.OrderStatusPaymentSent
		order.Payment.Method = entities.ToPaymentMethod(paymentMethod)
		if _, err := os.paymentUseCase.UpdatePayment(order.Payment); err != nil {
			return nil, err
		}
	}

	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrderConfirmed(order *entities.Order, paymentMethod string) (*entities.Order, error) {
	fmt.Printf("PaymentOrderConfirmed Order %d paid by method %s \n", order.ID, paymentMethod)
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NotPossibleWithoutItems, "PAY", order.ID)
	}

	if (order.Status != entities.OrderStatusPaymentError) && (order.Status != entities.OrderStatusPaymentSent) {
		return nil, fmt.Errorf("it is not possible to PAY the order: %d without PAYMENT_SENT or PAYMENT_ERROR", order.ID)
	}

	payment, err := os.paymentUseCase.GetPaymentById(order.Payment.ID)
	if err != nil {
		return nil, fmt.Errorf("error while getting payment id %v - %v", order.Payment.ID, err)
	}

	fmt.Printf("confimation payment %v\n", payment.ID)
	order.Status = entities.OrderStatusPaid
	payment.Status = entities.PaymentStatusPaid
	payment.ErrorReason = ""
	payment.Method = entities.ToPaymentMethod(paymentMethod)
	_, err = os.paymentUseCase.UpdatePayment(payment)
	if err != nil {
		return nil, fmt.Errorf("error while updating payment id %v - %v", payment.ID, err)
	}

	order.Payment = payment
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) PaymentOrderError(order *entities.Order, paymentMethod string, errorReason string) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NotPossibleWithoutItems, "PAY", order.ID)
	}

	if order.Status != entities.OrderStatusPaymentSent {
		return nil, fmt.Errorf("it is not possible to PAY the order: %d without PAYMENT_SENT", order.ID)
	}

	payment, err := os.paymentUseCase.GetPaymentById(order.Payment.ID)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusPaymentError
	payment.Status = entities.PaymentStatusError
	payment.Method = entities.ToPaymentMethod(paymentMethod)
	payment.ErrorReason = errorReason
	_, err = os.paymentUseCase.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}

	order.Payment = payment
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) DeliveredOrder(order *entities.Order) (*entities.Order, error) {
	if order.Payment.Status != entities.PaymentStatusPaid {
		return nil, fmt.Errorf(NotPossibleWithoutPayment, "DELIVERY", order.ID)
	}

	if order.Status != entities.OrderStatusSentForDelivery {
		return nil, fmt.Errorf("it is not possible to DELIVERY the order: %d without SENT FOR DELIVERY", order.ID)
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
		return nil, fmt.Errorf(NotPossibleWithoutItems, "PREPARE", order.ID)
	}

	if order.Status != entities.OrderStatusPaid {
		return nil, fmt.Errorf(NotPossibleWithoutPayment, "PREPARE", order.ID)
	}

	err := os.kitchenPublisher.PublishKitchen(order)
	if err != nil {
		return nil, err
	}

	order.Status = entities.OrderStatusKitchenPreparation
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) ReadyForDeliveryOrder(order *entities.Order) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NotPossibleWithoutPayment, "DELIVERY", order.ID)
	}

	if order.Payment.Status != entities.PaymentStatusPaid {
		return nil, fmt.Errorf(NotPossibleWithoutPayment, "DELIVERY", order.ID)
	}

	if order.Status != entities.OrderStatusKitchenPreparation {
		return nil, fmt.Errorf("it is not possible to DELIVERY the order: %d without PREPARE", order.ID)
	}

	order.Status = entities.OrderStatusReadyForDelivery
	return os.orderRepository.UpdateOrder(order)
}

func (os *OrderService) SentForDeliveryOrder(order *entities.Order) (*entities.Order, error) {
	if len(order.Items) == 0 {
		return nil, fmt.Errorf(NotPossibleWithoutItems, "DELIVERY", order.ID)
	}

	if order.Payment.Status != entities.PaymentStatusPaid {
		return nil, fmt.Errorf(NotPossibleWithoutPayment, "DELIVERY", order.ID)
	}

	if order.Status != entities.OrderStatusReadyForDelivery {
		return nil, fmt.Errorf("it is not possible to DELIVERY the order: %d without READY FOR DELIVERY", order.ID)
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
	if order.Status == entities.OrderStatusSentForDelivery || order.Status == entities.OrderStatusDelivered {
		return nil, fmt.Errorf("it is not possible to cancel the order: %d with status in [SentForDelivery, Delivered]", order.ID)
	}

	// TODO
	//if order.Payment.Status == entities.PaymentStatusPaid {
	//	payment, _ := os.paymentUseCase.GetPaymentById(order.Payment.ID)
	//	if err := os.paymentClient.Reverse(order); err != nil {
	//		order.Status = entities.OrderStatusPaymentError
	//		payment.Status = entities.PaymentStatusError
	//		return nil, err
	//	} else {
	//		payment.Status = entities.PaymentStatusReversed
	//		_, err = os.paymentUseCase.UpdatePayment(payment)
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		order.Payment = payment
	//	}
	//}

	order.Status = entities.OrderStatusCanceled
	return os.orderRepository.UpdateOrder(order)
}
