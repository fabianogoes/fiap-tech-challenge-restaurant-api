package dbo

import (
	"time"

	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"gorm.io/gorm"
)

// Order is a Database Object for order
type Order struct {
	gorm.Model
	CustomerID  uint
	Customer    Customer `gorm:"ForeignKey:CustomerID"`
	AttendantID uint
	Attendant   Attendant `gorm:"ForeignKey:AttendantID"`
	Date        time.Time
	Status      string
	PaymentID   uint
	Payment     Payment `gorm:"ForeignKey:PaymentID"`
	Amount      float64
	DeliveryID  uint
	Delivery    Delivery `gorm:"ForeignKey:DeliveryID"`
	Items       []*OrderItem
}

// ToEntity converts Order DBO to entity.Order
func (o *Order) ToEntity() *entity.Order {
	var items []*entity.OrderItem
	var itemsTotal int

	for _, item := range o.Items {
		items = append(items, item.ToEntity())
		itemsTotal += int(item.Quantity)
	}

	return &entity.Order{
		ID: o.ID,
		Customer: &entity.Customer{
			ID:   o.Customer.ID,
			Name: o.Customer.Name,
			CPF:  o.Customer.CPF,
		},
		Attendant: &entity.Attendant{
			ID:   o.Attendant.ID,
			Name: o.Attendant.Name,
		},
		Date:      o.Date,
		Status:    o.toOrderStatus(),
		Payment:   o.Payment.ToEntity(),
		Delivery:  o.Delivery.ToEntity(),
		Items:     items,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

// ToDBO converts entity.Order to Order DBO
func (o *Order) toOrderStatus() entity.OrderStatus {
	switch o.Status {
	case "STARTED":
		return entity.OrderStatusStarted
	case "ADDING_ITEMS":
		return entity.OrderStatusAddingItems
	case "CONFIRMED":
		return entity.OrderStatusConfirmed
	case "PAID":
		return entity.OrderStatusPaid
	case "PAYMENT_REVERSED":
		return entity.OrderStatusPaymentReversed
	case "IN_PREPARATION":
		return entity.OrderStatusInPreparation
	case "READY_FOR_DELIVERY":
		return entity.OrderStatusReadyForDelivery
	case "SENT_FOR_DELIVERY":
		return entity.OrderStatusSentForDelivery
	case "DELIVERED":
		return entity.OrderStatusDelivered
	case "CANCELED":
		return entity.OrderStatusCanceled
	default:
		return entity.OrderStatusStarted
	}
}

// ToDBO converts entity.Order to Order DBO
type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     Order
	ProductID uint
	Product   *Product
	Quantity  int
	UnitPrice float64
}

// ToEntity converts OrderItem DBO to entity.OrderItem
func (i *OrderItem) ToEntity() *entity.OrderItem {
	return &entity.OrderItem{
		ID:        i.ID,
		Product:   i.Product.ToEntity(),
		Quantity:  int(i.Quantity),
		UnitPrice: i.UnitPrice,
	}
}

// ToDBO converts entity.OrderItem to OrderItem DBO
func ToOrderItemDBO(i *entity.OrderItem) *OrderItem {
	return &OrderItem{
		Model: gorm.Model{
			ID:        i.ID,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
		},
		OrderID:   i.Order.ID,
		ProductID: i.Product.ID,
		Quantity:  i.Quantity,
		UnitPrice: i.UnitPrice,
	}
}
