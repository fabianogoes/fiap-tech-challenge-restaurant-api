package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/usecases"
	"github.com/fabianogoes/fiap-challenge/frameworks/rest/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestOrder_GetOrders(t *testing.T) {
	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("GetOrders").Return([]*entities.Order{domain.OrderStarted}, nil)

	orderUseCase := usecases.NewOrderService(
		orderRepository,
		new(domain.CustomerRepositoryMock),
		new(domain.AttendantRepositoryMock),
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		usecases.NewAttendantService(new(domain.AttendantRepositoryMock)),
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	setup := SetupTest()
	setup.GET("/orders", handler.GetOrders)
	request, err := http.NewRequest("GET", "/orders", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrders")
}

func TestOrder_GetOrdersNoContent(t *testing.T) {
	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("GetOrders").Return([]*entities.Order{}, nil)

	orderUseCase := usecases.NewOrderService(
		orderRepository,
		new(domain.CustomerRepositoryMock),
		new(domain.AttendantRepositoryMock),
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		usecases.NewAttendantService(new(domain.AttendantRepositoryMock)),
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	setup := SetupTest()
	setup.GET("/orders", handler.GetOrders)
	request, err := http.NewRequest("GET", "/orders", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrders")
}

func TestOrder_GetOrdersInternalServerError(t *testing.T) {
	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("GetOrders").Return(nil, errors.New("internal error"))

	orderUseCase := usecases.NewOrderService(
		orderRepository,
		new(domain.CustomerRepositoryMock),
		new(domain.AttendantRepositoryMock),
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		usecases.NewAttendantService(new(domain.AttendantRepositoryMock)),
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	setup := SetupTest()
	setup.GET("/orders", handler.GetOrders)
	request, err := http.NewRequest("GET", "/orders", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrders")
}

func TestOrder_GetOrderById(t *testing.T) {
	orderID := uint(1)
	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("GetOrderById", orderID).Return(domain.OrderStarted, nil)

	orderUseCase := usecases.NewOrderService(
		orderRepository,
		new(domain.CustomerRepositoryMock),
		new(domain.AttendantRepositoryMock),
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		usecases.NewAttendantService(new(domain.AttendantRepositoryMock)),
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	setup := SetupTest()
	setup.GET("/orders/:id", handler.GetOrderById)
	request, err := http.NewRequest("GET", fmt.Sprintf("/orders/%d", orderID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", orderID)
}

func TestOrder_GetOrderByIdBadRequest(t *testing.T) {
	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("GetOrderById", mock.Anything).Return(nil, errors.New("StatusBadRequest"))

	orderUseCase := usecases.NewOrderService(
		orderRepository,
		new(domain.CustomerRepositoryMock),
		new(domain.AttendantRepositoryMock),
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		usecases.NewAttendantService(new(domain.AttendantRepositoryMock)),
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	setup := SetupTest()
	setup.GET("/orders/:id", handler.GetOrderById)
	request, err := http.NewRequest("GET", "/orders/x", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertNotCalled(t, "GetOrderById", mock.Anything)
}

func TestOrder_GetOrderByIdNotFound(t *testing.T) {
	orderID := uint(1)
	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("GetOrderById", orderID).Return(nil, errors.New("StatusNotFound"))

	orderUseCase := usecases.NewOrderService(
		orderRepository,
		new(domain.CustomerRepositoryMock),
		new(domain.AttendantRepositoryMock),
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		usecases.NewAttendantService(new(domain.AttendantRepositoryMock)),
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	setup := SetupTest()
	setup.GET("/orders/:id", handler.GetOrderById)
	request, err := http.NewRequest("GET", fmt.Sprintf("/orders/%d", orderID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNotFound, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", orderID)
}

func TestOrder_StartOrder(t *testing.T) {
	orderUseCase := usecases.NewOrderService(
		new(domain.OrderRepositoryMock),
		new(domain.CustomerRepositoryMock),
		new(domain.AttendantRepositoryMock),
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		usecases.NewAttendantService(new(domain.AttendantRepositoryMock)),
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	setup := SetupTest()
	setup.POST("/orders/", handler.StartOrder)
	request, err := http.NewRequest("POST", "/orders/", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrder_StartOrderBadRequestAttendant(t *testing.T) {
	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("GetOrderById", mock.Anything).Return(domain.OrderStarted, nil)
	payload := dto.StartOrderRequest{CustomerCPF: "123", AttendantID: uint(2)}

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", payload.AttendantID).Return(nil, errors.New("attendant not found"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		new(domain.CustomerRepositoryMock),
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		usecases.NewCustomerService(new(domain.CustomerRepositoryMock)),
		attendantUseCase,
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/orders/", handler.StartOrder)
	request, err := http.NewRequest("POST", "/orders/", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	attendantRepository.AssertCalled(t, "GetAttendantById", payload.AttendantID)
}

func TestOrder_StartOrderBadRequestCustomer(t *testing.T) {
	orderRepository := new(domain.OrderRepositoryMock)
	payload := dto.StartOrderRequest{CustomerCPF: "123", AttendantID: uint(2)}

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", payload.AttendantID).Return(domain.AttendantSuccess, nil)

	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerByCPF", payload.CustomerCPF).Return(nil, errors.New("customer not found"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/orders/", handler.StartOrder)
	request, err := http.NewRequest("POST", "/orders/", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	attendantRepository.AssertCalled(t, "GetAttendantById", payload.AttendantID)
	customerRepository.AssertCalled(t, "GetCustomerByCPF", payload.CustomerCPF)
}

func TestOrder_StartOrderInternalServerError(t *testing.T) {
	payload := dto.StartOrderRequest{CustomerCPF: "123", AttendantID: domain.AttendantSuccess.ID}

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", payload.AttendantID).Return(domain.AttendantSuccess, nil)

	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerByCPF", payload.CustomerCPF).Return(domain.CustomerSuccess, nil)
	customerRepository.On("GetCustomerById", domain.AttendantSuccess.ID).Return(domain.CustomerSuccess, nil)

	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("CreateOrder", mock.Anything).Return(nil, errors.New("order service error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/orders/", handler.StartOrder)
	request, err := http.NewRequest("POST", "/orders/", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	attendantRepository.AssertCalled(t, "GetAttendantById", payload.AttendantID)
	customerRepository.AssertCalled(t, "GetCustomerByCPF", payload.CustomerCPF)
	customerRepository.AssertCalled(t, "GetCustomerById", domain.CustomerSuccess.ID)
	orderRepository.AssertCalled(t, "CreateOrder", mock.Anything)
}

func TestOrder_StartOrderSuccess(t *testing.T) {
	payload := dto.StartOrderRequest{CustomerCPF: "123", AttendantID: domain.AttendantSuccess.ID}

	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantRepository.On("GetAttendantById", payload.AttendantID).Return(domain.AttendantSuccess, nil)

	customerRepository := new(domain.CustomerRepositoryMock)
	customerRepository.On("GetCustomerByCPF", payload.CustomerCPF).Return(domain.CustomerSuccess, nil)
	customerRepository.On("GetCustomerById", domain.AttendantSuccess.ID).Return(domain.CustomerSuccess, nil)

	orderRepository := new(domain.OrderRepositoryMock)
	orderRepository.On("CreateOrder", mock.Anything).Return(domain.OrderStarted, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		usecases.NewProductService(new(domain.ProductRepositoryMock)),
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/orders/", handler.StartOrder)
	request, err := http.NewRequest("POST", "/orders/", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	attendantRepository.AssertCalled(t, "GetAttendantById", payload.AttendantID)
	customerRepository.AssertCalled(t, "GetCustomerByCPF", payload.CustomerCPF)
	customerRepository.AssertCalled(t, "GetCustomerById", domain.CustomerSuccess.ID)
	orderRepository.AssertCalled(t, "CreateOrder", mock.Anything)
}

func TestOrder_AddItemToOrderSuccess(t *testing.T) {
	payload := dto.AddItemToOrderRequest{ProductID: domain.ProductSuccess.ID, Quantity: 1}
	orderWithItem := domain.OrderStarted
	domain.OrderStarted.Items = []*entities.OrderItem{
		{
			ID:        uint(1),
			Order:     *domain.OrderStarted,
			Product:   &entities.Product{},
			UnitPrice: 1,
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	productRepository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, nil)
	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderWithItem, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/:id/item", handler.AddItemToOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item", domain.OrderStarted.ID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	productRepository.AssertCalled(t, "GetProductById", mock.Anything)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
}

func TestOrder_AddItemInternalServerError(t *testing.T) {
	payload := dto.AddItemToOrderRequest{ProductID: domain.ProductSuccess.ID, Quantity: 1}
	domain.OrderStarted.Items = []*entities.OrderItem{
		{
			ID:        uint(1),
			Order:     *domain.OrderStarted,
			Product:   &entities.Product{},
			UnitPrice: 1,
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	productRepository.On("GetProductById", mock.Anything).Return(domain.ProductSuccess, nil)
	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update order error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/:id/item", handler.AddItemToOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item", domain.OrderStarted.ID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	productRepository.AssertCalled(t, "GetProductById", mock.Anything)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
}

func TestOrder_AddItemBadRequestGetProduct(t *testing.T) {
	payload := dto.AddItemToOrderRequest{ProductID: domain.ProductSuccess.ID, Quantity: 1}
	orderWithItem := domain.OrderStarted
	domain.OrderStarted.Items = []*entities.OrderItem{
		{
			ID:        uint(1),
			Order:     *domain.OrderStarted,
			Product:   &entities.Product{},
			UnitPrice: 1,
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	productRepository.On("GetProductById", mock.Anything).Return(nil, errors.New("get product error"))
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderWithItem, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/:id/item", handler.AddItemToOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item", domain.OrderStarted.ID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	productRepository.AssertCalled(t, "GetProductById", mock.Anything)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_AddItemBadRequestGetOrder(t *testing.T) {
	payload := dto.AddItemToOrderRequest{ProductID: domain.ProductSuccess.ID, Quantity: 1}
	domain.OrderStarted.Items = []*entities.OrderItem{
		{
			ID:        uint(1),
			Order:     *domain.OrderStarted,
			Product:   &entities.Product{},
			UnitPrice: 1,
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", mock.Anything).Return(nil, errors.New("get order error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/:id/item", handler.AddItemToOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item", domain.OrderStarted.ID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_AddItemBadRequestPayload(t *testing.T) {
	domain.OrderStarted.Items = []*entities.OrderItem{
		{
			ID:        uint(1),
			Order:     *domain.OrderStarted,
			Product:   &entities.Product{},
			UnitPrice: 1,
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.POST("/:id/item", handler.AddItemToOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrder_AddItemBadRequestId(t *testing.T) {
	domain.OrderStarted.Items = []*entities.OrderItem{
		{
			ID:        uint(1),
			Order:     *domain.OrderStarted,
			Product:   &entities.Product{},
			UnitPrice: 1,
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.POST("/:id/item", handler.AddItemToOrder)
	request, err := http.NewRequest("POST", "/x/item", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrder_RemoveItemToOrderSuccess(t *testing.T) {
	payload := dto.AddItemToOrderRequest{ProductID: domain.ProductSuccess.ID, Quantity: 1}
	orderWithItem := domain.OrderStarted
	itemID := uint(1)
	orderWithItem.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderWithItem.Payment = &entities.Payment{Status: entities.PaymentStatusPending}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(orderWithItem, nil)
	orderRepository.On("GetOrderItemById", itemID).Return(domain.OrderItemSuccess, nil)
	orderRepository.On("RemoveItemFromOrder", itemID).Return(nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderWithItem, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/:id/item/:idItem", handler.RemoveItemFromOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item/%d", domain.OrderStarted.ID, itemID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
	orderRepository.AssertCalled(t, "GetOrderItemById", itemID)
	orderRepository.AssertCalled(t, "RemoveItemFromOrder", itemID)
}

func TestOrder_RemoveItemInternalServerError(t *testing.T) {
	payload := dto.AddItemToOrderRequest{ProductID: domain.ProductSuccess.ID, Quantity: 1}
	orderWithItem := domain.OrderStarted
	itemID := uint(1)
	orderWithItem.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderWithItem.Payment = &entities.Payment{Status: entities.PaymentStatusPending}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(orderWithItem, nil)
	orderRepository.On("GetOrderItemById", itemID).Return(domain.OrderItemSuccess, nil)
	orderRepository.On("RemoveItemFromOrder", itemID).Return(nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("remove item error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/:id/item/:idItem", handler.RemoveItemFromOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item/%d", domain.OrderStarted.ID, itemID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
	orderRepository.AssertCalled(t, "GetOrderItemById", itemID)
	orderRepository.AssertCalled(t, "RemoveItemFromOrder", itemID)
}

func TestOrder_RemoveItemBadRequestGetOrder(t *testing.T) {
	itemID := uint(1)

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(nil, errors.New("get order error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.POST("/:id/item/:idItem", handler.RemoveItemFromOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item/%d", domain.OrderStarted.ID, itemID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_RemoveItemBadRequestItemId(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.POST("/:id/item/:idItem", handler.RemoveItemFromOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/%d/item/x", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrder_RemoveItemBadRequestOrderId(t *testing.T) {
	orderWithItem := domain.OrderStarted
	itemID := uint(1)
	orderWithItem.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderWithItem.Payment = &entities.Payment{Status: entities.PaymentStatusPending}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.POST("/:id/item/:idItem", handler.RemoveItemFromOrder)
	request, err := http.NewRequest("POST", fmt.Sprintf("/x/item/%d", itemID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrder_ConfirmationOrderSuccess(t *testing.T) {
	orderConfirmed := domain.OrderStarted
	orderConfirmed.Status = entities.OrderStatusConfirmed
	orderConfirmed.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderConfirmed.Payment = &entities.Payment{Status: entities.PaymentStatusPending}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderConfirmed, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/confirmation", handler.ConfirmationOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/confirmation", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
}

func TestOrder_ConfirmationInternalServerError(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update order error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/confirmation", handler.ConfirmationOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/confirmation", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
}

func TestOrder_ConfirmationBadRequestGetOrder(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(nil, errors.New("get order error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/confirmation", handler.ConfirmationOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/confirmation", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_ConfirmationBadRequestId(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/confirmation", handler.ConfirmationOrder)
	request, err := http.NewRequest("PUT", "/x/confirmation", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

// TODO
//func TestOrder_PaymentOrderSuccess(t *testing.T) {
//	payload := dto.PaymentOrderRequest{PaymentMethod: entities.PaymentMethodCreditCard.ToString()}
//	orderConfirmed := domain.OrderStarted
//	orderConfirmed.Status = entities.OrderStatusConfirmed
//	orderConfirmed.Items = []*entities.OrderItem{domain.OrderItemSuccess}
//	orderConfirmed.Payment = &entities.Payment{Status: entities.PaymentStatusPending}
//
//	attendantRepository := new(domain.AttendantRepositoryMock)
//	customerRepository := new(domain.CustomerRepositoryMock)
//	orderRepository := new(domain.OrderRepositoryMock)
//	productRepository := new(domain.ProductRepositoryMock)
//
//	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
//	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderConfirmed, nil)
//
//	attendantUseCase := usecases.NewAttendantService(attendantRepository)
//	customerUseCase := usecases.NewCustomerService(customerRepository)
//	productUseCase := usecases.NewProductService(productRepository)
//	orderUseCase := usecases.NewOrderService(
//		orderRepository,
//		customerRepository,
//		attendantRepository,
//		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
//		new(domain.DeliveryClientMock),
//		new(domain.DeliveryRepositoryMock),
//		new(domain.KitchenPublisherMock),
//		new(domain.PaymentPublisherMock),
//	)
//
//	handler := NewOrderHandler(
//		orderUseCase,
//		customerUseCase,
//		attendantUseCase,
//		productUseCase,
//	)
//
//	jsonRequest, _ := json.Marshal(payload)
//	readerPayload := bytes.NewReader(jsonRequest)
//
//	setup := SetupTest()
//	setup.PUT("/:id/payment", handler.PaymentOrder)
//	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment", domain.OrderStarted.ID), readerPayload)
//	assert.NoError(t, err)
//
//	response := httptest.NewRecorder()
//	setup.ServeHTTP(response, request)
//	assert.Equal(t, http.StatusOK, response.Code)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, response)
//	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
//	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
//}

// TODO
//func TestOrder_PaymentInternalServerError(t *testing.T) {
//	payload := dto.PaymentOrderRequest{PaymentMethod: entities.PaymentMethodCreditCard.ToString()}
//	orderConfirmed := domain.OrderStarted
//	orderConfirmed.Status = entities.OrderStatusConfirmed
//	orderConfirmed.Items = []*entities.OrderItem{domain.OrderItemSuccess}
//	orderConfirmed.Payment = &entities.Payment{Status: entities.PaymentStatusPending}
//
//	attendantRepository := new(domain.AttendantRepositoryMock)
//	customerRepository := new(domain.CustomerRepositoryMock)
//	orderRepository := new(domain.OrderRepositoryMock)
//	productRepository := new(domain.ProductRepositoryMock)
//
//	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
//	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update order error"))
//
//	attendantUseCase := usecases.NewAttendantService(attendantRepository)
//	customerUseCase := usecases.NewCustomerService(customerRepository)
//	productUseCase := usecases.NewProductService(productRepository)
//	orderUseCase := usecases.NewOrderService(
//		orderRepository,
//		customerRepository,
//		attendantRepository,
//		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
//		new(domain.DeliveryClientMock),
//		new(domain.DeliveryRepositoryMock),
//		new(domain.KitchenPublisherMock),
//		new(domain.PaymentPublisherMock),
//	)
//
//	handler := NewOrderHandler(
//		orderUseCase,
//		customerUseCase,
//		attendantUseCase,
//		productUseCase,
//	)
//
//	jsonRequest, _ := json.Marshal(payload)
//	readerPayload := bytes.NewReader(jsonRequest)
//
//	setup := SetupTest()
//	setup.PUT("/:id/payment", handler.PaymentOrder)
//	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment", domain.OrderStarted.ID), readerPayload)
//	assert.NoError(t, err)
//
//	response := httptest.NewRecorder()
//	setup.ServeHTTP(response, request)
//	assert.Equal(t, http.StatusInternalServerError, response.Code)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, response)
//	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
//	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
//}

func TestOrder_PaymentBadRequestPayload(t *testing.T) {
	orderConfirmed := domain.OrderStarted
	orderConfirmed.Status = entities.OrderStatusConfirmed
	orderConfirmed.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderConfirmed.Payment = &entities.Payment{Status: entities.PaymentStatusPending}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/payment", handler.PaymentOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_PaymentBadRequestGetOrder(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(nil, errors.New("get order error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/payment", handler.PaymentOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_PaymentBadRequestId(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/payment", handler.PaymentOrder)
	request, err := http.NewRequest("PUT", "/x/payment", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrder_PaymentWebhookSuccess(t *testing.T) {
	payload := dto.PaymentWebhookRequest{Status: entities.PaymentStatusPaid.ToString(), PaymentMethod: entities.PaymentMethodCreditCard.ToString()}
	orderPaymentSent := domain.OrderStarted
	orderPaymentSent.Status = entities.OrderStatusPaymentSent
	orderPaymentSent.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderPaymentSent.Payment = domain.PaymentPaid

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(orderPaymentSent, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderPaymentSent, nil)
	paymentRepository.On("GetPaymentById", mock.Anything).Return(domain.PaymentPaid, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(domain.PaymentPaid, nil)

	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.PUT("/:id/payment/webhook", handler.PaymentWebhook)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment/webhook", domain.OrderStarted.ID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
	paymentRepository.AssertCalled(t, "GetPaymentById", mock.Anything)
	paymentRepository.AssertCalled(t, "UpdatePayment", mock.Anything)
}

func TestOrder_PaymentWebhookInternalServerErrorPaid(t *testing.T) {
	payload := dto.PaymentWebhookRequest{Status: entities.PaymentStatusPaid.ToString(), PaymentMethod: entities.PaymentMethodCreditCard.ToString()}
	orderPaymentSent := domain.OrderStarted
	orderPaymentSent.Status = entities.OrderStatusPaymentSent
	orderPaymentSent.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderPaymentSent.Payment = domain.PaymentPaid

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(orderPaymentSent, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update order error"))
	paymentRepository.On("GetPaymentById", mock.Anything).Return(domain.PaymentPaid, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(domain.PaymentPaid, nil)

	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.PUT("/:id/payment/webhook", handler.PaymentWebhook)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment/webhook", domain.OrderStarted.ID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
	paymentRepository.AssertCalled(t, "GetPaymentById", mock.Anything)
	paymentRepository.AssertCalled(t, "UpdatePayment", mock.Anything)
}

func TestOrder_PaymentWebhookInternalServerErrorPending(t *testing.T) {
	payload := dto.PaymentWebhookRequest{Status: entities.PaymentStatusPending.ToString(), PaymentMethod: entities.PaymentMethodCreditCard.ToString()}
	orderPaymentSent := domain.OrderStarted
	orderPaymentSent.Status = entities.OrderStatusPaymentSent
	orderPaymentSent.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderPaymentSent.Payment = domain.PaymentPaid

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(orderPaymentSent, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update order error"))
	paymentRepository.On("GetPaymentById", mock.Anything).Return(domain.PaymentPaid, nil)
	paymentRepository.On("UpdatePayment", mock.Anything).Return(domain.PaymentPaid, nil)

	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.PUT("/:id/payment/webhook", handler.PaymentWebhook)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment/webhook", domain.OrderStarted.ID), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
	paymentRepository.AssertCalled(t, "GetPaymentById", mock.Anything)
	paymentRepository.AssertCalled(t, "UpdatePayment", mock.Anything)
}

func TestOrder_PaymentWebhookBadRequestPayload(t *testing.T) {
	orderPaymentSent := domain.OrderStarted
	orderPaymentSent.Status = entities.OrderStatusPaymentSent
	orderPaymentSent.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderPaymentSent.Payment = domain.PaymentPaid

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(orderPaymentSent, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderPaymentSent, nil)

	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/payment/webhook", handler.PaymentWebhook)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment/webhook", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_PaymentWebhookBadRequestGetOrder(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(nil, errors.New("get order error"))

	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/payment/webhook", handler.PaymentWebhook)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/payment/webhook", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_PaymentWebhookBadRequestId(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)

	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/payment/webhook", handler.PaymentWebhook)
	request, err := http.NewRequest("PUT", "/x/payment/webhook", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

// TODO
//func TestOrder_InPreparationOrderSuccess(t *testing.T) {
//	orderPaid := domain.OrderStarted
//	orderPaid.Status = entities.OrderStatusPaid
//	orderPaid.Items = []*entities.OrderItem{domain.OrderItemSuccess}
//	orderPaid.Payment = &entities.Payment{Status: entities.PaymentStatusPending}
//
//	attendantRepository := new(domain.AttendantRepositoryMock)
//	customerRepository := new(domain.CustomerRepositoryMock)
//	orderRepository := new(domain.OrderRepositoryMock)
//	productRepository := new(domain.ProductRepositoryMock)
//
//	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
//	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderPaid, nil)
//
//	attendantUseCase := usecases.NewAttendantService(attendantRepository)
//	customerUseCase := usecases.NewCustomerService(customerRepository)
//	productUseCase := usecases.NewProductService(productRepository)
//	orderUseCase := usecases.NewOrderService(
//		orderRepository,
//		customerRepository,
//		attendantRepository,
//		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
//		new(domain.DeliveryClientMock),
//		new(domain.DeliveryRepositoryMock),
//		new(domain.KitchenPublisherMock),
//		new(domain.PaymentPublisherMock),
//	)
//
//	handler := NewOrderHandler(
//		orderUseCase,
//		customerUseCase,
//		attendantUseCase,
//		productUseCase,
//	)
//
//	setup := SetupTest()
//	setup.PUT("/:id/in-preparation", handler.InPreparationOrder)
//	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/in-preparation", domain.OrderStarted.ID), nil)
//	assert.NoError(t, err)
//
//	response := httptest.NewRecorder()
//	setup.ServeHTTP(response, request)
//	assert.Equal(t, http.StatusOK, response.Code)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, response)
//	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
//	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
//}

// TODO
//func TestOrder_InPreparationInternalServerError(t *testing.T) {
//	orderPaid := domain.OrderStarted
//	orderPaid.Status = entities.OrderStatusPaid
//	orderPaid.Items = []*entities.OrderItem{domain.OrderItemSuccess}
//	orderPaid.Payment = &entities.Payment{Status: entities.PaymentStatusPending}
//
//	attendantRepository := new(domain.AttendantRepositoryMock)
//	customerRepository := new(domain.CustomerRepositoryMock)
//	orderRepository := new(domain.OrderRepositoryMock)
//	productRepository := new(domain.ProductRepositoryMock)
//
//	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
//	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("update order error"))
//
//	attendantUseCase := usecases.NewAttendantService(attendantRepository)
//	customerUseCase := usecases.NewCustomerService(customerRepository)
//	productUseCase := usecases.NewProductService(productRepository)
//	orderUseCase := usecases.NewOrderService(
//		orderRepository,
//		customerRepository,
//		attendantRepository,
//		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
//		new(domain.DeliveryClientMock),
//		new(domain.DeliveryRepositoryMock),
//		new(domain.KitchenPublisherMock),
//		new(domain.PaymentPublisherMock),
//	)
//
//	handler := NewOrderHandler(
//		orderUseCase,
//		customerUseCase,
//		attendantUseCase,
//		productUseCase,
//	)
//
//	setup := SetupTest()
//	setup.PUT("/:id/in-preparation", handler.InPreparationOrder)
//	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/in-preparation", domain.OrderStarted.ID), nil)
//	assert.NoError(t, err)
//
//	response := httptest.NewRecorder()
//	setup.ServeHTTP(response, request)
//	assert.Equal(t, http.StatusInternalServerError, response.Code)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, response)
//	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
//	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
//}

func TestOrder_InPreparationBadRequestGetOrder(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(nil, errors.New("get order error"))

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/in-preparation", handler.InPreparationOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/in-preparation", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}

func TestOrder_InPreparationBadRequestId(t *testing.T) {
	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/in-preparation", handler.InPreparationOrder)
	request, err := http.NewRequest("PUT", "/x/in-preparation", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestOrder_ReadyForDeliveryOrderSuccess(t *testing.T) {
	orderInPreparation := domain.OrderStarted
	orderInPreparation.Status = entities.OrderStatusKitchenPreparation
	orderInPreparation.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderInPreparation.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderInPreparation, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/ready-for-delivery", handler.ReadyForDeliveryOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/ready-for-delivery", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
}

func TestOrder_SentForDeliveryOrderSuccess(t *testing.T) {
	orderReadyForDelivery := domain.OrderStarted
	orderReadyForDelivery.Status = entities.OrderStatusReadyForDelivery
	orderReadyForDelivery.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderReadyForDelivery.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	orderReadyForDelivery.Delivery = domain.DeliverySuccess

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderReadyForDelivery, nil)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(domain.DeliverySuccess, nil)
	deliveryRepository.On("UpdateDelivery", mock.Anything).Return(domain.DeliverySuccess, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.DeliveryClientMock),
		deliveryRepository,
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/sent-for-delivery", handler.SentForDeliveryOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/sent-for-delivery", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
	deliveryRepository.AssertCalled(t, "GetDeliveryById", mock.Anything)
	deliveryRepository.AssertCalled(t, "UpdateDelivery", mock.Anything)
}

func TestOrder_DeliveredOrderSuccess(t *testing.T) {
	orderSentForDelivery := domain.OrderStarted
	orderSentForDelivery.Status = entities.OrderStatusSentForDelivery
	orderSentForDelivery.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderSentForDelivery.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	orderSentForDelivery.Delivery = domain.DeliverySuccess

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	deliveryClient := new(domain.DeliveryClientMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderSentForDelivery, nil)
	deliveryRepository.On("GetDeliveryById", mock.Anything).Return(domain.DeliverySuccess, nil)
	deliveryRepository.On("UpdateDelivery", mock.Anything).Return(domain.DeliverySuccess, nil)
	deliveryClient.On("Deliver", mock.Anything).Return(nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		deliveryClient,
		deliveryRepository,
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/delivered", handler.DeliveredOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/delivered", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
	deliveryRepository.AssertCalled(t, "GetDeliveryById", mock.Anything)
	deliveryRepository.AssertCalled(t, "UpdateDelivery", mock.Anything)
	deliveryClient.AssertCalled(t, "Deliver", mock.Anything)
}

func TestOrder_CancelOrderNotPaidSuccess(t *testing.T) {
	orderConfirmed := domain.OrderStarted
	orderConfirmed.Status = entities.OrderStatusConfirmed
	orderConfirmed.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderConfirmed.Payment = &entities.Payment{Status: entities.PaymentStatusPending}
	orderConfirmed.Delivery = domain.DeliverySuccess

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)
	deliveryClient := new(domain.DeliveryClientMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderConfirmed, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		deliveryClient,
		deliveryRepository,
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/cancel", handler.CancelOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/cancel", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
}

// TODO
//func TestOrder_CancelOrderPaidSuccess(t *testing.T) {
//	orderPaid := domain.OrderStarted
//	orderPaid.Status = entities.OrderStatusPaid
//	orderPaid.Items = []*entities.OrderItem{domain.OrderItemSuccess}
//	orderPaid.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
//	orderPaid.Delivery = domain.DeliverySuccess
//
//	attendantRepository := new(domain.AttendantRepositoryMock)
//	customerRepository := new(domain.CustomerRepositoryMock)
//	orderRepository := new(domain.OrderRepositoryMock)
//	productRepository := new(domain.ProductRepositoryMock)
//	deliveryRepository := new(domain.DeliveryRepositoryMock)
//	paymentRepository := new(domain.PaymentRepositoryMock)
//	deliveryClient := new(domain.DeliveryClientMock)
//
//	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
//	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderPaid, nil)
//	paymentRepository.On("GetPaymentById", mock.Anything).Return(orderPaid.Payment, nil)
//	paymentRepository.On("UpdatePayment", mock.Anything).Return(orderPaid.Payment, nil)
//
//	attendantUseCase := usecases.NewAttendantService(attendantRepository)
//	customerUseCase := usecases.NewCustomerService(customerRepository)
//	productUseCase := usecases.NewProductService(productRepository)
//	paymentUseCase := usecases.NewPaymentService(paymentRepository)
//	orderUseCase := usecases.NewOrderService(
//		orderRepository,
//		customerRepository,
//		attendantRepository,
//		paymentUseCase,
//		deliveryClient,
//		deliveryRepository,
//		new(domain.KitchenPublisherMock),
//		new(domain.PaymentPublisherMock),
//	)
//
//	handler := NewOrderHandler(
//		orderUseCase,
//		customerUseCase,
//		attendantUseCase,
//		productUseCase,
//	)
//
//	setup := SetupTest()
//	setup.PUT("/:id/cancel", handler.CancelOrder)
//	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/cancel", domain.OrderStarted.ID), nil)
//	assert.NoError(t, err)
//
//	response := httptest.NewRecorder()
//	setup.ServeHTTP(response, request)
//	assert.Equal(t, http.StatusOK, response.Code)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, response)
//	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
//	orderRepository.AssertCalled(t, "UpdateOrder", mock.Anything, mock.Anything, mock.Anything)
//	paymentRepository.AssertCalled(t, "GetPaymentById", mock.Anything)
//	paymentRepository.AssertCalled(t, "UpdatePayment", mock.Anything)
//}

// TODO
//func TestOrder_CancelOrderReverseFailed(t *testing.T) {
//	orderPaid := domain.OrderStarted
//	orderPaid.Status = entities.OrderStatusPaid
//	orderPaid.Items = []*entities.OrderItem{domain.OrderItemSuccess}
//	orderPaid.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
//	orderPaid.Delivery = domain.DeliverySuccess
//
//	attendantRepository := new(domain.AttendantRepositoryMock)
//	customerRepository := new(domain.CustomerRepositoryMock)
//	orderRepository := new(domain.OrderRepositoryMock)
//	productRepository := new(domain.ProductRepositoryMock)
//	deliveryRepository := new(domain.DeliveryRepositoryMock)
//	paymentRepository := new(domain.PaymentRepositoryMock)
//	deliveryClient := new(domain.DeliveryClientMock)
//
//	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)
//	orderRepository.On("UpdateOrder", mock.Anything, mock.Anything, mock.Anything).Return(orderPaid, nil)
//	paymentRepository.On("GetPaymentById", mock.Anything).Return(orderPaid.Payment, nil)
//	paymentRepository.On("UpdatePayment", mock.Anything).Return(orderPaid.Payment, nil)
//
//	attendantUseCase := usecases.NewAttendantService(attendantRepository)
//	customerUseCase := usecases.NewCustomerService(customerRepository)
//	productUseCase := usecases.NewProductService(productRepository)
//	paymentUseCase := usecases.NewPaymentService(paymentRepository)
//	orderUseCase := usecases.NewOrderService(
//		orderRepository,
//		customerRepository,
//		attendantRepository,
//		paymentUseCase,
//		deliveryClient,
//		deliveryRepository,
//		new(domain.KitchenPublisherMock),
//		new(domain.PaymentPublisherMock),
//	)
//
//	handler := NewOrderHandler(
//		orderUseCase,
//		customerUseCase,
//		attendantUseCase,
//		productUseCase,
//	)
//
//	setup := SetupTest()
//	setup.PUT("/:id/cancel", handler.CancelOrder)
//	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/cancel", domain.OrderStarted.ID), nil)
//	assert.NoError(t, err)
//
//	response := httptest.NewRecorder()
//	setup.ServeHTTP(response, request)
//	assert.Equal(t, http.StatusInternalServerError, response.Code)
//
//	assert.NoError(t, err)
//	assert.NotNil(t, response)
//	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
//	paymentRepository.AssertCalled(t, "GetPaymentById", mock.Anything)
//}

func TestOrder_CancelOrderDeliveredFailed(t *testing.T) {
	orderPaid := domain.OrderStarted
	orderPaid.Status = entities.OrderStatusDelivered
	orderPaid.Items = []*entities.OrderItem{domain.OrderItemSuccess}
	orderPaid.Payment = &entities.Payment{Status: entities.PaymentStatusPaid}
	orderPaid.Delivery = domain.DeliverySuccess

	attendantRepository := new(domain.AttendantRepositoryMock)
	customerRepository := new(domain.CustomerRepositoryMock)
	orderRepository := new(domain.OrderRepositoryMock)
	productRepository := new(domain.ProductRepositoryMock)
	deliveryRepository := new(domain.DeliveryRepositoryMock)
	paymentRepository := new(domain.PaymentRepositoryMock)
	deliveryClient := new(domain.DeliveryClientMock)

	orderRepository.On("GetOrderById", domain.OrderStarted.ID).Return(domain.OrderStarted, nil)

	attendantUseCase := usecases.NewAttendantService(attendantRepository)
	customerUseCase := usecases.NewCustomerService(customerRepository)
	productUseCase := usecases.NewProductService(productRepository)
	paymentUseCase := usecases.NewPaymentService(paymentRepository)
	orderUseCase := usecases.NewOrderService(
		orderRepository,
		customerRepository,
		attendantRepository,
		paymentUseCase,
		deliveryClient,
		deliveryRepository,
		new(domain.KitchenPublisherMock),
		new(domain.PaymentPublisherMock),
	)

	handler := NewOrderHandler(
		orderUseCase,
		customerUseCase,
		attendantUseCase,
		productUseCase,
	)

	setup := SetupTest()
	setup.PUT("/:id/cancel", handler.CancelOrder)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/%d/cancel", domain.OrderStarted.ID), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	orderRepository.AssertCalled(t, "GetOrderById", domain.OrderStarted.ID)
}
