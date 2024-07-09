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

func TestBDDOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Order BDD Suite")
}

var _ = Describe("Order", func() {
	item := entities.OrderItem{
		ID:        uint(101010),
		Product:   domain.ProductSuccess,
		Quantity:  10,
		UnitPrice: 10_00,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	order := entities.Order{
		ID:        uint(101010),
		Customer:  domain.OrderStarted.Customer,
		Attendant: domain.OrderStarted.Attendant,
		Payment: &entities.Payment{
			ID:        1,
			Order:     entities.Order{},
			Date:      time.Now(),
			Method:    entities.PaymentMethodCreditCard,
			Status:    entities.PaymentStatusPending,
			Value:     10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Date:      time.Now(),
		Status:    entities.OrderStatusAddingItems,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Items:     []*entities.OrderItem{&item},
	}

	Context("order confirmed", func() {
		orderConfirmed := order
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

		response, err := service.ConfirmationOrder(&order)
		It("has no error on confirmation", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be Confirmed", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusConfirmed))
		})
	})

	Context("order sent to payment", func() {
		orderPaymentSent := order
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

		response, err := service.PaymentOrder(&order, entities.PaymentMethodCreditCard.ToString())
		It("has no error on PaymentOrder", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be PaymentSent", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusPaymentSent))
		})
	})

	Context("order paid", func() {
		orderPaid := order
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

		orderPaymentSent := order
		orderPaymentSent.Status = entities.OrderStatusPaymentSent
		orderPaymentSent.Payment = &paymentPaid
		response, err := service.PaymentOrderConfirmed(&orderPaymentSent, entities.PaymentMethodCreditCard.ToString())
		It("has no error on PaymentOrderConfirmed", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be Paid", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusPaid))
		})
	})

	Context("order payment error", func() {
		orderPaymentError := order
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

		orderPaymentSent := order
		orderPaymentSent.Status = entities.OrderStatusPaymentSent
		orderPaymentSent.Payment = &paymentPaid
		response, err := service.PaymentOrderError(&orderPaymentSent, entities.PaymentMethodCreditCard.ToString(), "invalid card")
		It("has no error on PaymentOrderError", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be PaymentError", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusPaymentError))
		})
	})

	Context("in preparation", func() {
		orderInPreparation := order
		orderInPreparation.Status = entities.OrderStatusInPreparation
		repository := new(domain.OrderRepositoryMock)
		repository.On("UpdateOrder", mock.Anything).Return(&orderInPreparation, nil)
		kitchenClient := new(domain.KitchenClientMock)
		kitchenClient.On("Preparation", mock.Anything).Return(nil)
		service := NewOrderService(
			repository,
			new(domain.CustomerRepositoryMock),
			new(domain.AttendantRepositoryMock),
			NewPaymentService(new(domain.PaymentRepositoryMock)),
			new(domain.PaymentClientMock),
			new(domain.DeliveryClientMock),
			new(domain.DeliveryRepositoryMock),
			kitchenClient,
		)

		orderPaid := order
		orderPaid.Status = entities.OrderStatusPaid
		response, err := service.InPreparationOrder(&orderPaid)
		It("has no error on in preparation", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be InPreparation", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusInPreparation))
		})
	})

	Context("ready for delivery", func() {
		orderInPreparation := order
		orderInPreparation.Status = entities.OrderStatusInPreparation

		orderReadyForDelivery := order
		orderReadyForDelivery.Status = entities.OrderStatusReadyForDelivery

		repository := new(domain.OrderRepositoryMock)
		repository.On("UpdateOrder", mock.Anything).Return(&orderReadyForDelivery, nil)
		kitchenClient := new(domain.KitchenClientMock)
		kitchenClient.On("ReadyDelivery", mock.Anything).Return(nil)

		deliveryRepository := new(domain.DeliveryRepositoryMock)
		service := NewOrderService(
			repository,
			new(domain.CustomerRepositoryMock),
			new(domain.AttendantRepositoryMock),
			NewPaymentService(new(domain.PaymentRepositoryMock)),
			new(domain.PaymentClientMock),
			new(domain.DeliveryClientMock),
			deliveryRepository,
			kitchenClient,
		)

		orderInPreparation.Payment.Status = entities.PaymentStatusPaid
		response, err := service.ReadyForDeliveryOrder(&orderInPreparation)
		It("has no error on ReadyForDeliveryOrder", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be ReadyForDelivery", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusReadyForDelivery))
		})

	})

	Context("set for delivery", func() {
		orderReadyForDelivery := order
		orderReadyForDelivery.Status = entities.OrderStatusReadyForDelivery
		deliveryPending := entities.Delivery{
			ID:        uint(2020),
			Order:     entities.Order{},
			Status:    entities.DeliveryStatusPending,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		deliverySent := deliveryPending
		deliverySent.Status = entities.DeliveryStatusSent
		orderReadyForDelivery.Delivery = &deliveryPending

		orderSentForDelivery := order
		orderSentForDelivery.Status = entities.OrderStatusSentForDelivery
		orderSentForDelivery.Delivery = &deliverySent

		repository := new(domain.OrderRepositoryMock)
		repository.On("UpdateOrder", mock.Anything).Return(&orderSentForDelivery, nil)
		kitchenClient := new(domain.KitchenClientMock)
		kitchenClient.On("ReadyDelivery", mock.Anything).Return(nil)

		deliveryRepository := new(domain.DeliveryRepositoryMock)
		deliveryRepository.On("GetDeliveryById", mock.Anything).Return(&deliveryPending, nil)
		deliveryRepository.On("UpdateDelivery", mock.Anything).Return(&deliverySent, nil)
		service := NewOrderService(
			repository,
			new(domain.CustomerRepositoryMock),
			new(domain.AttendantRepositoryMock),
			NewPaymentService(new(domain.PaymentRepositoryMock)),
			new(domain.PaymentClientMock),
			new(domain.DeliveryClientMock),
			deliveryRepository,
			kitchenClient,
		)

		response, err := service.SentForDeliveryOrder(&orderReadyForDelivery)
		It("has no error on SentForDeliveryOrder", func() {
			Expect(err).Should(BeNil())
		})

		It("has order not be nil", func() {
			Expect(response).ShouldNot(BeNil())
		})

		It("has order status be SentForDelivery", func() {
			Expect(response.Status).Should(Equal(entities.OrderStatusSentForDelivery))
		})

		It("has delivery status be Sent", func() {
			Expect(response.Delivery.Status).Should(Equal(entities.DeliveryStatusSent))
		})
	})

})
