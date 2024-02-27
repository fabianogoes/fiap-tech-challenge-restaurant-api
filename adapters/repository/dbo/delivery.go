package dbo

import (
	"github.com/fabianogoes/fiap-challenge/entities"
	"time"

	"gorm.io/gorm"
)

// Delivery is a Database Object for delivery
type Delivery struct {
	gorm.Model
	Date   time.Time
	Status string
}

// ToEntity converts Delivery DBO to entities.Delivery
func (d *Delivery) ToEntity() *entities.Delivery {
	return &entities.Delivery{
		ID:     d.ID,
		Date:   d.Date,
		Status: d.toDeliveryStatus(),
	}
}

// ToDBO converts entities.Delivery to Delivery DBO
func ToDeliveryDBO(d *entities.Delivery) Delivery {
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

// toDeliveryStatus converts string to entities.DeliveryStatus
func (d *Delivery) toDeliveryStatus() entities.DeliveryStatus {
	switch d.Status {
	case "PENDING":
		return entities.DeliveryStatusPending
	case "SENT":
		return entities.DeliveryStatusSent
	case "DELIVERED":
		return entities.DeliveryStatusDelivered
	case "CANCELED":
		return entities.DeliveryStatusCanceled
	default:
		return entities.DeliveryStatusPending
	}
}
