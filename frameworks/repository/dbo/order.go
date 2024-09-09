package dbo

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"time"

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

// ToEntity converts Order DBO to entities.Order
func (o *Order) ToEntity() *entities.Order {
	var items []*entities.OrderItem
	var itemsTotal int

	for _, item := range o.Items {
		items = append(items, item.ToEntity())
		itemsTotal += int(item.Quantity)
	}

	return &entities.Order{
		ID: o.ID,
		Customer: &entities.Customer{
			ID:   o.Customer.ID,
			Name: o.Customer.Name,
			CPF:  o.Customer.CPF,
		},
		Attendant: &entities.Attendant{
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

// ToDBO converts entities.Order to Order DBO
func (o *Order) toOrderStatus() entities.OrderStatus {
	switch o.Status {
	case "STARTED":
		return entities.OrderStatusStarted
	case "ADDING_ITEMS":
		return entities.OrderStatusAddingItems
	case "CONFIRMED":
		return entities.OrderStatusConfirmed
	case "PAID":
		return entities.OrderStatusPaid
	case "PAYMENT_SENT":
		return entities.OrderStatusPaymentSent
	case "PAYMENT_ERROR":
		return entities.OrderStatusPaymentError
	case "PAYMENT_REVERSED":
		return entities.OrderStatusPaymentReversed
	case "KITCHEN_WAITING":
		return entities.OrderStatusKitchenWaiting
	case "KITCHEN_PREPARATION":
		return entities.OrderStatusKitchenPreparation
	case "KITCHEN_READY":
		return entities.OrderStatusKitchenReady
	case "KITCHEN_CANCELED":
		return entities.OrderStatusKitchenCanceled
	case "READY_FOR_DELIVERY":
		return entities.OrderStatusReadyForDelivery
	case "SENT_FOR_DELIVERY":
		return entities.OrderStatusSentForDelivery
	case "DELIVERED":
		return entities.OrderStatusDelivered
	case "CANCELED":
		return entities.OrderStatusCanceled
	default:
		return entities.OrderStatusStarted
	}
}

// OrderItem ToDBO converts entities.Order to Order DBO
type OrderItem struct {
	gorm.Model
	OrderID   uint
	Order     Order
	ProductID uint
	Product   *Product
	Quantity  int
	UnitPrice float64
}

// ToEntity converts OrderItem DBO to entities.OrderItem
func (i *OrderItem) ToEntity() *entities.OrderItem {
	return &entities.OrderItem{
		ID:        i.ID,
		Product:   i.Product.ToEntity(),
		Quantity:  int(i.Quantity),
		UnitPrice: i.UnitPrice,
	}
}

// ToOrderItemDBO ToDBO converts entities.OrderItem to OrderItem DBO
func ToOrderItemDBO(i *entities.OrderItem) *OrderItem {
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
