package dbo

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

// Delivery is a Database Object for delivery
type Delivery struct {
	gorm.Model
	Date   time.Time
	Status string
}

// ToEntity converts Delivery DBO to domain.Delivery
func (d *Delivery) ToEntity() *domain.Delivery {
	return &domain.Delivery{
		ID:     d.ID,
		Date:   d.Date,
		Status: d.toDeliveryStatus(),
	}
}

// ToDBO converts domain.Delivery to Delivery DBO
func ToDeliveryDBO(d *domain.Delivery) Delivery {
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

// toDeliveryStatus converts string to domain.DeliveryStatus
func (d *Delivery) toDeliveryStatus() domain.DeliveryStatus {
	switch d.Status {
	case "PENDING":
		return domain.DeliveryStatusPending
	case "SENT":
		return domain.DeliveryStatusSent
	case "DELIVERED":
		return domain.DeliveryStatusDelivered
	case "CANCELED":
		return domain.DeliveryStatusCanceled
	default:
		return domain.DeliveryStatusPending
	}
}
