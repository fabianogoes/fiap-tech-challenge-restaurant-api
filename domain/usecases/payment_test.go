package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var paymentIDSuccess = uint(1)
var PaymentPending = &entities.Payment{
	ID:        paymentIDSuccess,
	Order:     *OrderStarted,
	Method:    entities.PaymentMethodCreditCard,
	Status:    entities.PaymentStatusPending,
	Value:     100_00,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestPaymentService_GetPaymentById(t *testing.T) {
	repository := new(domain.PaymentRepositoryMock)
	repository.On("GetPaymentById", mock.Anything).Return(PaymentPending, nil)

	service := NewPaymentService(repository)
	paymentResponse, err := service.GetPaymentById(paymentIDSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, paymentResponse)
}

func TestPaymentService_UpdatePayment(t *testing.T) {
	repository := new(domain.PaymentRepositoryMock)
	repository.On("UpdatePayment", mock.Anything).Return(PaymentPending, nil)

	service := NewPaymentService(repository)
	paymentResponse, err := service.UpdatePayment(PaymentPending)
	assert.NoError(t, err)
	assert.NotNil(t, paymentResponse)
}
