package rest

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/usecases"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Router(t *testing.T) {
	config := entities.NewConfig()
	customerRepository := new(domain.CustomerRepositoryMock)
	customerHandler := NewCustomerHandler(usecases.NewCustomerService(customerRepository), config)
	attendantRepository := new(domain.AttendantRepositoryMock)
	attendantHandler := NewAttendantHandler(usecases.NewAttendantService(attendantRepository))
	productRepository := new(domain.ProductRepositoryMock)
	productHandler := NewProductHandler(usecases.NewProductService(productRepository))
	orderService := usecases.NewOrderService(
		new(domain.OrderRepositoryMock),
		customerRepository,
		attendantRepository,
		usecases.NewPaymentService(new(domain.PaymentRepositoryMock)),
		new(domain.PaymentClientMock),
		new(domain.DeliveryClientMock),
		new(domain.DeliveryRepositoryMock),
		new(domain.KitchenClientMock),
	)
	orderHandler := NewOrderHandler(
		orderService,
		usecases.NewCustomerService(customerRepository),
		usecases.NewAttendantService(attendantRepository),
		usecases.NewProductService(productRepository),
	)

	router, err := NewRouter(
		customerHandler,
		attendantHandler,
		productHandler,
		orderHandler,
	)
	assert.NoError(t, err)
	assert.NotNil(t, router)
}
