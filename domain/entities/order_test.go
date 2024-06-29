package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Order(t *testing.T) {
	order, err := NewOrder(&Customer{}, &Attendant{})
	assert.NoError(t, err)
	assert.NotNil(t, order)

	product := &Product{Price: 1}
	order.AddItem(product, 1)
	assert.Equal(t, product.Price, order.Amount())

	assert.Equal(t, 1, order.ItemsQuantity())
}

func Test_OrderStatusToString(t *testing.T) {
	assert.Equal(t, "STARTED", OrderStatusStarted.ToString())
	assert.Equal(t, "ADDING_ITEMS", OrderStatusAddingItems.ToString())
	assert.Equal(t, "CONFIRMED", OrderStatusConfirmed.ToString())
	assert.Equal(t, "PAID", OrderStatusPaid.ToString())
	assert.Equal(t, "PAYMENT_SENT", OrderStatusPaymentSent.ToString())
	assert.Equal(t, "PAYMENT_REVERSED", OrderStatusPaymentReversed.ToString())
	assert.Equal(t, "PAYMENT_ERROR", OrderStatusPaymentError.ToString())
	assert.Equal(t, "IN_PREPARATION", OrderStatusInPreparation.ToString())
	assert.Equal(t, "READY_FOR_DELIVERY", OrderStatusReadyForDelivery.ToString())
	assert.Equal(t, "SENT_FOR_DELIVERY", OrderStatusSentForDelivery.ToString())
	assert.Equal(t, "DELIVERED", OrderStatusDelivered.ToString())
	assert.Equal(t, "DELIVERY_ERROR", OrderStatusDeliveryError.ToString())
	assert.Equal(t, "CANCELED", OrderStatusCanceled.ToString())
	assert.Equal(t, "UNKNOWN", OrderStatusUnknown.ToString())
}

func Test_ToOrderStatus(t *testing.T) {
	assert.Equal(t, OrderStatusStarted, ToOrderStatus("STARTED"))
	assert.Equal(t, OrderStatusAddingItems, ToOrderStatus("ADDING_ITEMS"))
	assert.Equal(t, OrderStatusConfirmed, ToOrderStatus("CONFIRMED"))
	assert.Equal(t, OrderStatusPaid, ToOrderStatus("PAID"))
	assert.Equal(t, OrderStatusPaymentSent, ToOrderStatus("PAYMENT_SENT"))
	assert.Equal(t, OrderStatusPaymentReversed, ToOrderStatus("PAYMENT_REVERSED"))
	assert.Equal(t, OrderStatusPaymentError, ToOrderStatus("PAYMENT_ERROR"))
	assert.Equal(t, OrderStatusInPreparation, ToOrderStatus("IN_PREPARATION"))
	assert.Equal(t, OrderStatusReadyForDelivery, ToOrderStatus("READY_FOR_DELIVERY"))
	assert.Equal(t, OrderStatusSentForDelivery, ToOrderStatus("SENT_FOR_DELIVERY"))
	assert.Equal(t, OrderStatusDelivered, ToOrderStatus("DELIVERED"))
	assert.Equal(t, OrderStatusDeliveryError, ToOrderStatus("DELIVERY_ERROR"))
	assert.Equal(t, OrderStatusCanceled, ToOrderStatus("CANCELED"))
	assert.Equal(t, OrderStatusUnknown, ToOrderStatus("UNKNOWN"))
}
