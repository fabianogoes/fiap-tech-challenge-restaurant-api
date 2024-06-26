package usecases

import (
	"errors"
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var orderIDSuccess = uint(1)
var orderItemIDSuccess = uint(1)
var OrderItemSuccess = &entities.OrderItem{
	ID:        orderItemIDSuccess,
	Product:   ProductSuccess,
	Quantity:  10,
	UnitPrice: 10_00,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
var OrderStarted = &entities.Order{
	ID:        orderIDSuccess,
	Customer:  CustomerSuccess,
	Attendant: AttendantSuccess,
	Date:      time.Now(),
	Status:    entities.OrderStatusStarted,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestOrderService_StartOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("CreateOrder", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	order, err := service.StartOrder(CustomerSuccess.ID, AttendantSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_GetOrderById(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderById", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	order, err := service.GetOrderById(orderIDSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_GetOrders(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrders").Return([]*entities.Order{OrderStarted}, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	order, err := service.GetOrders()
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_AddItemToOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	order, err := service.AddItemToOrder(OrderStarted, ProductSuccess, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_RemoveItemFromOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	orderRequest := OrderStarted
	orderRequest.Payment = &entities.Payment{
		ID:     ProductSuccess.ID,
		Status: entities.PaymentStatusPending,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_RemoveItemFromOrderPaid(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Reverse", mock.Anything).Return(nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	orderRequest := OrderStarted
	orderRequest.Payment = &entities.Payment{
		ID:     ProductSuccess.ID,
		Status: entities.PaymentStatusPaid,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_RemoveItemFromOrderPaidReverseError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Reverse", mock.Anything).Return(errors.New("reverse error"))

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	orderRequest := OrderStarted
	orderRequest.Payment = &entities.Payment{
		ID:     ProductSuccess.ID,
		Status: entities.PaymentStatusPaid,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_ConfirmationOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	order, err := service.ConfirmationOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_PaymentOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Status = entities.OrderStatusConfirmed
	order, err := service.PaymentOrder(OrderStarted, "CREDIT_CARD")
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_PaymentOrderConfirmed(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Status = entities.OrderStatusPaymentSent
	OrderStarted.Payment = PaymentPending
	order, err := service.PaymentOrderConfirmed(OrderStarted, "CREDIT_CARD")
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_PaymentOrderError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Status = entities.OrderStatusPaymentSent
	OrderStarted.Payment = PaymentPending
	order, err := service.PaymentOrderError(OrderStarted, "CREDIT_CARD", "Error test")
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_DeliveredOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	delivery := &entities.Delivery{ID: 1}
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(delivery, nil)
	deliveryRepository.On("UpdateDelivery", mock.Anything).Return(delivery, nil)

	deliveryClient := new(domain.DeliveryClientMock)
	deliveryClient.On("Deliver", mock.Anything).Return(nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		deliveryClient,
		deliveryRepository,
		new(domain.KitchenClientMock),
	)

	PaymentPending.Status = entities.PaymentStatusPaid
	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Status = entities.OrderStatusSentForDelivery
	OrderStarted.Payment = PaymentPending
	OrderStarted.Delivery = delivery
	order, err := service.DeliveredOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_InPreparationOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	kitchenClient.On("Preparation", mock.Anything).Return(nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		kitchenClient,
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Status = entities.OrderStatusPaid
	order, err := service.InPreparationOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_ReadyForDeliveryOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	kitchenClient.On("ReadyDelivery", mock.Anything).Return(nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		kitchenClient,
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Status = entities.OrderStatusInPreparation
	OrderStarted.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	order, err := service.ReadyForDeliveryOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_SentForDeliveryOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	kitchenClient.On("ReadyDelivery", mock.Anything).Return(nil)

	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(&entities.Delivery{}, nil)
	deliveryRepository.On("UpdateDelivery", mock.Anything).Return(&entities.Delivery{}, nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		deliveryRepository,
		kitchenClient,
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Status = entities.OrderStatusReadyForDelivery
	OrderStarted.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	order, err := service.SentForDeliveryOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_CancelOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)

	paymentClient := new(domain.PaymentClientMock)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		deliveryRepository,
		kitchenClient,
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Payment = &entities.Payment{ID: 1}
	OrderStarted.Status = entities.OrderStatusStarted
	order, err := service.CancelOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_CancelOrderPaid(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(&entities.Payment{}, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)

	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Reverse", mock.Anything).Return(nil)

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		deliveryRepository,
		kitchenClient,
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Payment = &entities.Payment{ID: 1, Status: entities.PaymentStatusPaid}
	OrderStarted.Status = entities.OrderStatusStarted
	order, err := service.CancelOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_CancelOrderPaidReverseError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)

	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Reverse", mock.Anything).Return(errors.New("reverse error"))

	service := NewOrderService(
		repository,
		customerRepository,
		attendantRepository,
		paymentService,
		paymentClient,
		new(domain.DeliveryClientMock),
		deliveryRepository,
		kitchenClient,
	)

	OrderStarted.Items = []*entities.OrderItem{OrderItemSuccess}
	OrderStarted.Payment = &entities.Payment{ID: 1, Status: entities.PaymentStatusPaid}
	OrderStarted.Status = entities.OrderStatusStarted
	order, err := service.CancelOrder(OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}
