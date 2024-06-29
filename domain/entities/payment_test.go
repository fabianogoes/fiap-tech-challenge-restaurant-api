package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PaymentStatusToString(t *testing.T) {
	assert.Equal(t, "PENDING", PaymentStatusPending.ToString())
	assert.Equal(t, "PAID", PaymentStatusPaid.ToString())
	assert.Equal(t, "REVERSED", PaymentStatusReversed.ToString())
	assert.Equal(t, "CANCELED", PaymentStatusCanceled.ToString())
	assert.Equal(t, "ERROR", PaymentStatusError.ToString())
	assert.Equal(t, "NONE", PaymentStatusNone.ToString())
	assert.Equal(t, "UNKNOWN", PaymentStatusUnknown.ToString())
}

func Test_PaymentStatusToPaymentStatus(t *testing.T) {
	assert.Equal(t, PaymentStatusPending, ToPaymentStatus("PENDING"))
	assert.Equal(t, PaymentStatusPaid, ToPaymentStatus("PAID"))
	assert.Equal(t, PaymentStatusReversed, ToPaymentStatus("REVERSED"))
	assert.Equal(t, PaymentStatusCanceled, ToPaymentStatus("CANCELED"))
	assert.Equal(t, PaymentStatusError, ToPaymentStatus("ERROR"))
	assert.Equal(t, PaymentStatusNone, ToPaymentStatus("NONE"))
	assert.Equal(t, PaymentStatusUnknown, ToPaymentStatus("UNKNOWN"))
}

func Test_PaymentMethodToString(t *testing.T) {
	assert.Equal(t, "CREDIT_CARD", PaymentMethodCreditCard.ToString())
	assert.Equal(t, "DEBIT_CARD", PaymentMethodDebitCard.ToString())
	assert.Equal(t, "MONEY", PaymentMethodMoney.ToString())
	assert.Equal(t, "PIX", PaymentMethodPIX.ToString())
	assert.Equal(t, "NONE", PaymentMethodNone.ToString())
}

func Test_PaymentToPaymentMethod(t *testing.T) {
	assert.Equal(t, PaymentMethodCreditCard, ToPaymentMethod("CREDIT_CARD"))
	assert.Equal(t, PaymentMethodDebitCard, ToPaymentMethod("DEBIT_CARD"))
	assert.Equal(t, PaymentMethodMoney, ToPaymentMethod("MONEY"))
	assert.Equal(t, PaymentMethodPIX, ToPaymentMethod("PIX"))
	assert.Equal(t, PaymentMethodNone, ToPaymentMethod("NONE"))
}
