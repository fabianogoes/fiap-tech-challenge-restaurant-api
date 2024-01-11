package dbo

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

// Delivery is a Database Object for delivery
type Delivery struct {
	gorm.Model
	OrderID uint
	Order   Order
	Date    time.Time
	Status  string
}

// ToEntity converts Delivery DBO to entity.Delivery
func (d *Delivery) ToEntity() *entity.Delivery {
	return &entity.Delivery{
		ID:     d.ID,
		Date:   d.Date,
		Status: d.toDeliveryStatus(),
	}
}

// ToDBO converts entity.Delivery to Delivery DBO
func ToDeliveryDBO(d *entity.Delivery) Delivery {
	return Delivery{
		Model: gorm.Model{
			ID:        d.ID,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Date:   d.Date,
		Status: d.Status.ToString(),
	}
}

// toDeliveryStatus converts string to entity.DeliveryStatus
func (d *Delivery) toDeliveryStatus() entity.DeliveryStatus {
	switch d.Status {
	case "PENDING":
		return entity.DeliveryStatusPending
	case "SENT":
		return entity.DeliveryStatusSent
	case "DELIVERED":
		return entity.DeliveryStatusDelivered
	case "CANCELED":
		return entity.DeliveryStatusCanceled
	default:
		return entity.DeliveryStatusPending
	}
}
