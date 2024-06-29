package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_DeliveryToString(t *testing.T) {
	assert.Equal(t, "PENDING", DeliveryStatusPending.ToString())
	assert.Equal(t, "SENT", DeliveryStatusSent.ToString())
	assert.Equal(t, "DELIVERED", DeliveryStatusDelivered.ToString())
	assert.Equal(t, "CANCELED", DeliveryStatusCanceled.ToString())
	assert.Equal(t, "ERROR", DeliveryStatusError.ToString())
	assert.Equal(t, "NONE", DeliveryStatusNone.ToString())
}

func Test_DeliveryToDeliveryStatus(t *testing.T) {
	assert.Equal(t, DeliveryStatusPending, ToDeliveryStatus("PENDING"))
	assert.Equal(t, DeliveryStatusSent, ToDeliveryStatus("SENT"))
	assert.Equal(t, DeliveryStatusDelivered, ToDeliveryStatus("DELIVERED"))
	assert.Equal(t, DeliveryStatusCanceled, ToDeliveryStatus("CANCELED"))
	assert.Equal(t, DeliveryStatusError, ToDeliveryStatus("ERROR"))
	assert.Equal(t, DeliveryStatusNone, ToDeliveryStatus("NONE"))
}

func Test_Delivery(t *testing.T) {
	delivery := Delivery{
		ID:        1,
		Date:      time.Now(),
		Status:    DeliveryStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	assert.NotNil(t, delivery)
}
