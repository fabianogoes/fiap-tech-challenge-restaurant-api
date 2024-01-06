package repository

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID    uint
	Customer      Customer `gorm:"ForeignKey:CustomerID"`
	AttendantID   uint
	Attendant     Attendant `gorm:"ForeignKey:AttendantID"`
	Date          time.Time
	Status        string
	PaymentStatus string
	Amount        float64
	Items         []*OrderItem
}

func mapOrderStatus(status string) domain.OrderStatus {
	switch status {
	case "STARTED":
		return domain.OrderStatusStarted
	case "CONFIRMED":
		return domain.OrderStatusConfirmed
	case "IN_PREPARATION":
		return domain.OrderStatusInPreparation
	case "PAID":
		return domain.OrderStatusPaid
	case "DELIVERED":
		return domain.OrderStatusDelivered
	case "CANCELED":
		return domain.OrderStatusCanceled
	default:
		return domain.OrderStatusStarted
	}
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     Order
	ProductID uint
	Product   Product
	Quantity  int64
	UnitPrice float64
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (or *OrderRepository) StartOrder(
	customerID uint,
	attendantID uint,
	orderStatus string,
	paymentStatus string,
) (*domain.Order, error) {
	order := &Order{
		CustomerID:    customerID,
		AttendantID:   attendantID,
		Date:          time.Now(),
		Status:        orderStatus,
		PaymentStatus: paymentStatus,
		Amount:        0,
		Items:         []*OrderItem{},
	}

	if err := or.db.Create(order).Error; err != nil {
		return nil, err
	}

	return &domain.Order{
		ID: order.ID,
		Customer: &domain.Customer{
			ID:   order.Customer.ID,
			Name: order.Customer.Name,
			CPF:  order.Customer.CPF,
		},
		Attendant: &domain.Attendant{
			ID:   order.Attendant.ID,
			Name: order.Attendant.Name,
		},
		Date:          order.Date,
		Status:        mapOrderStatus(order.Status),
		PaymentStatus: order.PaymentStatus,
		Amount:        order.Amount,
		Items:         []*domain.OrderItem{},
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}, nil
}

// func (or *OrderRepository) StartOrder(customerID int, attendantID int) (*domain.Order, error)
// 	// order := &Order{
// 	// 	CustomerID:    customerID,
// 	// 	AttendantID:   attendantID,
// 	// 	Date:          time.Now(),
// 	// 	Status:        "started",
// 	// 	PaymentStatus: "pending",
// 	// 	Amount:        0,
// 	// 	Items:         []*OrderItem{},
// 	// }

// 	// if err := or.db.Create(order).Error; err != nil {
// 	// 	return nil, err
// 	// }

// 	return &order, nil
// }
