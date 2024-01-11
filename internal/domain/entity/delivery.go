package entity

import "time"

type DeliveryStatus int

const (
	DeliveryStatusPending DeliveryStatus = iota
	DeliveryStatusSent
	DeliveryStatusDelivered
	DeliveryStatusCanceled
)

func (ds DeliveryStatus) ToString() string {
	return [...]string{"PENDING", "SENT", "DELIVERED", "CANCELED"}[ds]
}

func (ds DeliveryStatus) ToDeliveryStatus(status string) DeliveryStatus {
	switch status {
	case "PENDING":
		return DeliveryStatusPending
	case "SENT":
		return DeliveryStatusSent
	case "DELIVERED":
		return DeliveryStatusDelivered
	case "CANCELED":
		return DeliveryStatusCanceled
	default:
		return DeliveryStatusPending
	}
}

type Delivery struct {
	ID        uint
	Order     Order
	Date      time.Time
	Status    DeliveryStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
