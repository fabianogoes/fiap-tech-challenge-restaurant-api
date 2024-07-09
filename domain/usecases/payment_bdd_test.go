package usecases

import (
	"errors"
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBDDPayment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payment BDD Suite")
}

var _ = Describe("Payment", func() {
	item := &entities.OrderItem{
		ID:        uint(101010),
		Product:   domain.ProductSuccess,
		Quantity:  10,
		UnitPrice: 10_00,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	order := &entities.Order{
		ID:        uint(101010),
		Customer:  domain.OrderStarted.Customer,
		Attendant: domain.OrderStarted.Attendant,
		Date:      time.Now(),
		Status:    entities.OrderStatusAddingItems,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Items:     []*entities.OrderItem{item},
	}

	Context("initially", func() {
		orderConfirmed := *order
		orderConfirmed.Status = entities.OrderStatusConfirmed
		repository := new(domain.OrderRepositoryMock)
		repository.On("UpdateOrder", mock.Anything).Return(&orderConfirmed, nil)
		service := NewOrderService(
			repository,
			new(domain.CustomerRepositoryMock),
			new(domain.AttendantRepositoryMock),
			NewPaymentService(new(domain.PaymentRepositoryMock)),
			new(domain.PaymentClientMock),
			new(domain.DeliveryClientMock),
			new(domain.DeliveryRepositoryMock),
			new(domain.KitchenClientMock),
		)

		response, err := service.ConfirmationOrder(order)
		It("has no error on get order", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be Confirmed", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusConfirmed))
		})
	})

	Context("payment sent", func() {
		orderPaymentSent := *order
		orderPaymentSent.Status = entities.OrderStatusPaymentSent

		repository := new(domain.OrderRepositoryMock)
		repository.On("UpdateOrder", mock.Anything).Return(&orderPaymentSent, nil)
		paymentClient := new(domain.PaymentClientMock)
		paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)
		service := NewOrderService(
			repository,
			new(domain.CustomerRepositoryMock),
			new(domain.AttendantRepositoryMock),
			NewPaymentService(new(domain.PaymentRepositoryMock)),
			paymentClient,
			new(domain.DeliveryClientMock),
			new(domain.DeliveryRepositoryMock),
			new(domain.KitchenClientMock),
		)

		response, err := service.PaymentOrder(order, entities.PaymentMethodCreditCard.ToString())
		It("has no error on get order", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be Confirmed", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusPaymentSent))
		})
	})

	Context("payment confirmed", func() {
		orderPaid := *order
		orderPaid.Status = entities.OrderStatusPaid

		repository := new(domain.OrderRepositoryMock)
		repository.On("UpdateOrder", mock.Anything).Return(&orderPaid, nil)
		paymentClient := new(domain.PaymentClientMock)
		paymentClient.On("Pay", mock.Anything, mock.Anything).Return(nil)
		paymentRepository := new(domain.PaymentRepositoryMock)

		paymentPending := entities.Payment{
			ID:        1,
			Order:     entities.Order{},
			Date:      time.Now(),
			Method:    entities.PaymentMethodCreditCard,
			Status:    entities.PaymentStatusPending,
			Value:     10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		paymentPaid := paymentPending
		paymentPaid.Status = entities.PaymentStatusPaid
		paymentRepository.On("GetPaymentById", mock.Anything).Return(&paymentPending, nil)
		paymentRepository.On("UpdatePayment", mock.Anything).Return(&paymentPaid, nil)
		service := NewOrderService(
			repository,
			new(domain.CustomerRepositoryMock),
			new(domain.AttendantRepositoryMock),
			NewPaymentService(paymentRepository),
			paymentClient,
			new(domain.DeliveryClientMock),
			new(domain.DeliveryRepositoryMock),
			new(domain.KitchenClientMock),
		)

		orderPaymentSent := *order
		orderPaymentSent.Status = entities.OrderStatusPaymentSent
		orderPaymentSent.Payment = &paymentPaid
		response, err := service.PaymentOrderConfirmed(&orderPaymentSent, entities.PaymentMethodCreditCard.ToString())
		It("has no error on get order", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be Confirmed", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusPaid))
		})
	})

	Context("payment error", func() {
		orderPaymentError := *order
		orderPaymentError.Status = entities.OrderStatusPaymentError

		repository := new(domain.OrderRepositoryMock)
		repository.On("UpdateOrder", mock.Anything).Return(&orderPaymentError, nil)
		paymentClient := new(domain.PaymentClientMock)
		paymentClient.On("Pay", mock.Anything, mock.Anything).Return(errors.New("payment error"))
		paymentRepository := new(domain.PaymentRepositoryMock)

		paymentPending := entities.Payment{
			ID:        1,
			Order:     entities.Order{},
			Date:      time.Now(),
			Method:    entities.PaymentMethodCreditCard,
			Status:    entities.PaymentStatusPending,
			Value:     10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		paymentPaid := paymentPending
		paymentPaid.Status = entities.PaymentStatusPaid
		paymentRepository.On("GetPaymentById", mock.Anything).Return(&paymentPending, nil)
		paymentRepository.On("UpdatePayment", mock.Anything).Return(&paymentPaid, nil)
		service := NewOrderService(
			repository,
			new(domain.CustomerRepositoryMock),
			new(domain.AttendantRepositoryMock),
			NewPaymentService(paymentRepository),
			paymentClient,
			new(domain.DeliveryClientMock),
			new(domain.DeliveryRepositoryMock),
			new(domain.KitchenClientMock),
		)

		orderPaymentSent := *order
		orderPaymentSent.Status = entities.OrderStatusPaymentSent
		orderPaymentSent.Payment = &paymentPaid
		response, err := service.PaymentOrderError(&orderPaymentSent, entities.PaymentMethodCreditCard.ToString(), "invalid card")
		It("has no error on get order", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be Confirmed", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusPaymentError))
		})
	})
})
