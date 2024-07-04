package usecases

import (
	"errors"
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestOrderService_StartOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("CreateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order, err := service.StartOrder(domain.CustomerSuccess.ID, domain.AttendantSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_StartOrderError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("CreateOrder", mock.Anything).Return(nil, errors.New("create order error"))

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

	order, err := service.StartOrder(domain.CustomerSuccess.ID, domain.AttendantSuccess.ID)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_StartOrderGetAttendantError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(nil, errors.New("get attendant error"))

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("CreateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order, err := service.StartOrder(domain.CustomerSuccess.ID, domain.AttendantSuccess.ID)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_StartOrderGetCustomerError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(nil, errors.New("get customer error"))

	attendantRepository := new(domain.AttendantRepositoryMock)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("CreateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order, err := service.StartOrder(domain.CustomerSuccess.ID, domain.AttendantSuccess.ID)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_GetOrderById(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderById", mock.Anything).Return(domain.OrderStarted, nil)

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

	order, err := service.GetOrderById(domain.OrderItemSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_GetOrders(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrders").Return([]*entities.Order{domain.OrderStarted}, nil)

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
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order, err := service.AddItemToOrder(domain.OrderStarted, domain.ProductSuccess, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_RemoveItemFromOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(domain.OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	orderRequest := domain.OrderStarted
	orderRequest.Payment = &entities.Payment{
		ID:     domain.ProductSuccess.ID,
		Status: entities.PaymentStatusPending,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_RemoveItemFromOrderError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(domain.OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(errors.New("remove item error"))
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	orderRequest := domain.OrderStarted
	orderRequest.Payment = &entities.Payment{
		ID:     domain.ProductSuccess.ID,
		Status: entities.PaymentStatusPending,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_RemoveItemFromOrderSentForDeliveryError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(domain.OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	orderRequest := domain.OrderStarted
	orderRequest.Status = entities.OrderStatusSentForDelivery
	orderRequest.Payment = &entities.Payment{
		ID:     domain.ProductSuccess.ID,
		Status: entities.PaymentStatusPending,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_RemoveItemFromOrderGetItemError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	attendantRepository := new(domain.AttendantRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(nil, errors.New("get item error"))

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

	orderRequest := domain.OrderStarted
	orderRequest.Payment = &entities.Payment{
		ID:     domain.ProductSuccess.ID,
		Status: entities.PaymentStatusPending,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_RemoveItemFromOrderPaid(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(domain.OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	orderRequest := domain.OrderStarted
	orderRequest.Status = entities.OrderStatusPaid
	orderRequest.Payment = &entities.Payment{
		ID:     domain.ProductSuccess.ID,
		Status: entities.PaymentStatusPaid,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_RemoveItemFromOrderPaidUpdatePaymentError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(nil, errors.New("update payment error"))
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(domain.OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	orderRequest := domain.OrderStarted
	orderRequest.Status = entities.OrderStatusPaid
	orderRequest.Payment = &entities.Payment{
		ID:     domain.ProductSuccess.ID,
		Status: entities.PaymentStatusPaid,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_RemoveItemFromOrderPaidReverseError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("GetOrderItemById", mock.Anything).Return(domain.OrderItemSuccess, nil)
	repository.On("RemoveItemFromOrder", mock.Anything).Return(nil)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	orderRequest := domain.OrderStarted
	orderRequest.Payment = &entities.Payment{
		ID:     domain.ProductSuccess.ID,
		Status: entities.PaymentStatusPaid,
	}

	order, err := service.RemoveItemFromOrder(orderRequest, 1)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_ConfirmationOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order, err := service.ConfirmationOrder(domain.OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_ConfirmationOrderItemsEmptyError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order, err := service.ConfirmationOrder(domain.OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_PaymentOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Status = entities.OrderStatusConfirmed
	order, err := service.PaymentOrder(domain.OrderStarted, "CREDIT_CARD")
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_PaymentOrderPayError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(errors.New("payment error"))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusConfirmed
	order, err := service.PaymentOrder(order, "CREDIT_CARD")
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, order.Status, entities.OrderStatusPaymentError)
}

func TestOrderService_PaymentOrderOrderNotConfirmedError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Status = entities.OrderStatusStarted
	order, err := service.PaymentOrder(domain.OrderStarted, "CREDIT_CARD")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderItemsEmptyError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{}
	domain.OrderStarted.Status = entities.OrderStatusConfirmed
	order, err := service.PaymentOrder(domain.OrderStarted, "CREDIT_CARD")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderConfirmed(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Status = entities.OrderStatusPaymentSent
	domain.OrderStarted.Payment = PaymentPending
	order, err := service.PaymentOrderConfirmed(domain.OrderStarted, "CREDIT_CARD")
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_PaymentOrderConfirmedUpdateError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(nil, errors.New("update payment error"))

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusPaymentSent
	order.Payment = PaymentPending
	order, err := service.PaymentOrderConfirmed(order, "CREDIT_CARD")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderConfirmedGetPaymentError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(nil, errors.New("get payment error"))
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusPaymentSent
	order.Payment = PaymentPending
	order, err := service.PaymentOrderConfirmed(order, "CREDIT_CARD")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderConfirmedNotSent(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusStarted
	order.Payment = PaymentPending
	order, err := service.PaymentOrderConfirmed(order, "CREDIT_CARD")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderConfirmedItemsEmptyError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)
	paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{}
	order.Status = entities.OrderStatusPaymentSent
	order.Payment = PaymentPending
	orderResponse, err := service.PaymentOrderConfirmed(order, "CREDIT_CARD")
	assert.Error(t, err)
	assert.Nil(t, orderResponse)
}

func TestOrderService_PaymentOrderError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Status = entities.OrderStatusPaymentSent
	domain.OrderStarted.Payment = PaymentPending
	order, err := service.PaymentOrderError(domain.OrderStarted, "CREDIT_CARD", "Error test")
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_PaymentOrderErrorUpdateError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	attendantRepository := new(domain.AttendantRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(nil, errors.New("update payment error"))

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)
	repository := new(domain.OrderRepositoryMock)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusPaymentSent
	order.Payment = PaymentPending
	order, err := service.PaymentOrderError(order, "CREDIT_CARD", "Error test")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderErrorGetPaymentError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(nil, errors.New("get payment error"))
	paymentRepository.On("UpdatePayment", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusPaymentSent
	order.Payment = PaymentPending
	order, err := service.PaymentOrderError(order, "CREDIT_CARD", "Error test")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderErrorNotSent(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusStarted
	order.Payment = PaymentPending
	order, err := service.PaymentOrderError(order, "CREDIT_CARD", "Error test")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_PaymentOrderErrorItemsEmpty(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{}
	order.Status = entities.OrderStatusPaymentSent
	order.Payment = PaymentPending
	order, err := service.PaymentOrderError(order, "CREDIT_CARD", "Error test")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_DeliveredOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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
	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusSentForDelivery
	order.Payment = PaymentPending
	order.Delivery = delivery
	response, err := service.DeliveredOrder(order)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrderService_DeliveredOrderNotPaid(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	PaymentPending.Status = entities.PaymentStatusPending
	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusSentForDelivery
	order.Payment = PaymentPending
	order.Delivery = delivery
	response, err := service.DeliveredOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_DeliveredOrderUpdateError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

	delivery := &entities.Delivery{ID: 1}
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(delivery, nil)
	deliveryRepository.On("UpdateDelivery", mock.Anything).Return(nil, errors.New("delivery update error"))

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
	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusSentForDelivery
	order.Payment = PaymentPending
	order.Delivery = delivery
	response, err := service.DeliveredOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_DeliveredOrderClientError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

	delivery := &entities.Delivery{ID: 1}
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(delivery, nil)
	deliveryRepository.On("UpdateDelivery", mock.Anything).Return(delivery, nil)

	deliveryClient := new(domain.DeliveryClientMock)
	deliveryClient.On("Deliver", mock.Anything).Return(errors.New("delivery client error"))

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
	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusSentForDelivery
	order.Payment = PaymentPending
	order.Delivery = delivery
	response, err := service.DeliveredOrder(order)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, entities.OrderStatusDeliveryError, order.Status)
}

func TestOrderService_DeliveredOrderGetError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	attendantRepository := new(domain.AttendantRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)

	delivery := &entities.Delivery{ID: 1}
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(nil, errors.New("get delivery error"))
	deliveryClient := new(domain.DeliveryClientMock)

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
	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusSentForDelivery
	order.Payment = PaymentPending
	order.Delivery = delivery
	response, err := service.DeliveredOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_DeliveredOrderNotSentForDelivery(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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
	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusStarted
	order.Payment = PaymentPending
	order.Delivery = delivery
	response, err := service.DeliveredOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_DeliveredOrderItemsEmptyError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)

	paymentService := NewPaymentService(paymentRepository)
	paymentClient := new(domain.PaymentClientMock)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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
	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{}
	order.Status = entities.OrderStatusSentForDelivery
	order.Payment = PaymentPending
	order.Delivery = delivery
	response, err := service.DeliveredOrder(domain.OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrderService_InPreparationOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusPaid
	response, err := service.InPreparationOrder(order)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrderService_InPreparationOrderClientError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	kitchenClient.On("Preparation", mock.Anything).Return(errors.New("client error"))

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusPaid
	response, err := service.InPreparationOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_InPreparationOrderNotPaid(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusStarted
	response, err := service.InPreparationOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_InPreparationOrderItemsEmptyError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{}
	order.Status = entities.OrderStatusPaid
	response, err := service.InPreparationOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_ReadyForDeliveryOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusInPreparation
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.ReadyForDeliveryOrder(order)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrderService_ReadyForDeliveryOrderClientError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

	kitchenClient := new(domain.KitchenClientMock)
	kitchenClient.On("ReadyDelivery", mock.Anything).Return(errors.New("client error"))

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusInPreparation
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.ReadyForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_ReadyForDeliveryOrderNotInPreparationError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusStarted
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.ReadyForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_ReadyForDeliveryOrderNotPaidError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusInPreparation
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPending}
	response, err := service.ReadyForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_ReadyForDeliveryOrderItemsEmptyError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{}
	order.Status = entities.OrderStatusInPreparation
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.ReadyForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_SentForDeliveryOrderNotPaid(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusReadyForDelivery
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPending}
	response, err := service.SentForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_SentForDeliveryOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusReadyForDelivery
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.SentForDeliveryOrder(order)
	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrderService_SentForDeliveryOrderUpdateError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	attendantRepository := new(domain.AttendantRepositoryMock)
	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))
	repository := new(domain.OrderRepositoryMock)
	kitchenClient := new(domain.KitchenClientMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(&entities.Delivery{}, nil)
	deliveryRepository.On("UpdateDelivery", mock.Anything).Return(nil, errors.New("update error"))

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusReadyForDelivery
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.SentForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_SentForDeliveryOrderGetError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	attendantRepository := new(domain.AttendantRepositoryMock)
	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))
	repository := new(domain.OrderRepositoryMock)
	kitchenClient := new(domain.KitchenClientMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(nil, errors.New("get delivery error"))
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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusReadyForDelivery
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.SentForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_SentForDeliveryOrderNotReadyForDeliveryError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Status = entities.OrderStatusStarted
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.SentForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_SentForDeliveryOrderItemsEmptyError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentService := NewPaymentService(new(domain.PaymentRepositoryMock))

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{}
	order.Status = entities.OrderStatusReadyForDelivery
	order.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	response, err := service.SentForDeliveryOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_CancelOrder(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Payment = &entities.Payment{ID: 1}
	domain.OrderStarted.Status = entities.OrderStatusStarted
	order, err := service.CancelOrder(domain.OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_CancelOrderUpdatePaymentError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(domain.PaymentPaid, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(nil, errors.New("update payment error"))
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	order := domain.OrderStarted
	order.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	order.Payment = &entities.Payment{ID: 1, Status: entities.PaymentStatusPaid}
	order.Status = entities.OrderStatusStarted
	response, err := service.CancelOrder(order)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestOrderService_CancelOrderSentForDelivery(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Payment = &entities.Payment{ID: 1}
	domain.OrderStarted.Status = entities.OrderStatusSentForDelivery
	order, err := service.CancelOrder(domain.OrderStarted)
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrderService_CancelOrderPaid(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(&entities.Payment{}, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Payment = &entities.Payment{ID: 1, Status: entities.PaymentStatusPaid}
	domain.OrderStarted.Status = entities.OrderStatusStarted
	order, err := service.CancelOrder(domain.OrderStarted)
	assert.NoError(t, err)
	assert.NotNil(t, order)
}

func TestOrderService_CancelOrderPaidReverseError(t *testing.T) {
	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	paymentRepository := new(domain.PaymentRepositoryMock)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(&entities.Payment{}, nil)
	paymentService := NewPaymentService(paymentRepository)

	repository := new(domain.OrderRepositoryMock)
	repository.On("UpdateOrder", mock.Anything).Return(domain.OrderStarted, nil)

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

	domain.OrderStarted.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	domain.OrderStarted.Payment = &entities.Payment{ID: 1, Status: entities.PaymentStatusPaid}
	domain.OrderStarted.Status = entities.OrderStatusStarted
	order, err := service.CancelOrder(domain.OrderStarted)
	assert.Error(t, err)
	assert.Nil(t, order)
}
