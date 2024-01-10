package entity

import "time"

type DeliveryStatus int

const (
	DeliveryStatusPending DeliveryStatus = iota
	DeliveryStatusSent
	DeliveryStatusDelivered
	DeliveryStatusCanceled
)

type Delivery struct {
	ID        int64
	Order     Order
	Date      time.Time
	Status    DeliveryStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
