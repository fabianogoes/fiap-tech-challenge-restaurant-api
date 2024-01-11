package entity

import "time"

type Order struct {
	ID        uint
	Customer  *Customer
	Attendant *Attendant
	Date      time.Time
	Status    OrderStatus
	Payment   *Payment
	Items     []*OrderItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrder(customer *Customer, attendant *Attendant) (*Order, error) {
	return &Order{
		Customer:  customer,
		Attendant: attendant,
		Date:      time.Now(),
		Status:    OrderStatusStarted,
		Items:     []*OrderItem{},
		Payment: &Payment{
			Status: PaymentStatusPending,
			Method: PaymentMethodNone,
		},
	}, nil
}

func (o *Order) Amount() float64 {
	var amount float64

	for _, item := range o.Items {
		amount += item.UnitPrice * float64(item.Quantity)
	}

	return amount
}

func (o *Order) AddItem(product *Product, quantity int) {
	o.Items = append(o.Items, &OrderItem{
		Order:     *o,
		Product:   product,
		Quantity:  quantity,
		UnitPrice: product.Price,
	})
}

func (o *Order) ItemsQuantity() int {
	var itemsTotal int

	for _, item := range o.Items {
		itemsTotal += item.Quantity
	}

	return itemsTotal
}

type OrderItem struct {
	ID        uint
	Order     Order
	Product   *Product
	Quantity  int
	UnitPrice float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderStatus int

const (
	OrderStatusStarted OrderStatus = iota
	OrderStatusAddingItems
	OrderStatusConfirmed
	OrderStatusPaid
	OrderStatusPaymentReversed
	OrderStatusPaymentError
	OrderStatusInPreparation
	OrderStatusReadyForDelivery
	OrderStatusSentForDelivery
	OrderStatusDelivered
	OrderStatusDeliveryError
	OrderStatusCanceled
)

func (os OrderStatus) ToString() string {
	return [...]string{
		"STARTED",
		"ADDING_ITEMS",
		"CONFIRMED",
		"PAID",
		"PAYMENT_REVERSED",
		"PAYMENT_ERROR",
		"IN_PREPARATION",
		"READY_FOR_DELIVERY",
		"SENT_FOR_DELIVERY",
		"DELIVERED",
		"DELIVERY_ERROR",
		"CANCELED",
	}[os]
}

func (os OrderStatus) ToOrderStatus(status string) OrderStatus {
	switch status {
	case "STARTED":
		return OrderStatusStarted
	case "CONFIRMED":
		return OrderStatusConfirmed
	case "IN_PREPARATION":
		return OrderStatusInPreparation
	case "READY_FOR_DELIVERY":
		return OrderStatusReadyForDelivery
	case "PAID":
		return OrderStatusPaid
	case "PAYMENT_REVERSED":
		return OrderStatusPaymentReversed
	case "PAYMENT_ERROR":
		return OrderStatusPaymentError
	case "SENT_FOR_DELIVERY":
		return OrderStatusSentForDelivery
	case "DELIVERED":
		return OrderStatusDelivered
	case "DELIVERY_ERROR":
		return OrderStatusDeliveryError
	case "CANCELED":
		return OrderStatusCanceled
	default:
		return OrderStatusStarted
	}
}
