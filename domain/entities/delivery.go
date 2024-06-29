package entities

import "time"

type DeliveryStatus int

const (
	DeliveryStatusPending DeliveryStatus = iota
	DeliveryStatusSent
	DeliveryStatusDelivered
	DeliveryStatusCanceled
	DeliveryStatusError
	DeliveryStatusNone
)

func (ds DeliveryStatus) ToString() string {
	return [...]string{"PENDING", "SENT", "DELIVERED", "CANCELED", "ERROR", "NONE"}[ds]
}

func ToDeliveryStatus(status string) DeliveryStatus {
	switch status {
	case "PENDING":
		return DeliveryStatusPending
	case "SENT":
		return DeliveryStatusSent
	case "DELIVERED":
		return DeliveryStatusDelivered
	case "CANCELED":
		return DeliveryStatusCanceled
	case "ERROR":
		return DeliveryStatusError
	default:
		return DeliveryStatusNone
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
