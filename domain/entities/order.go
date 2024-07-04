package entities

import "time"

type Order struct {
	ID        uint
	Customer  *Customer
	Attendant *Attendant
	Date      time.Time
	Status    OrderStatus
	Payment   *Payment
	Delivery  *Delivery
	Items     []*OrderItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewOrder(customer *Customer, attendant *Attendant) *Order {
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
		Delivery: &Delivery{
			Status: DeliveryStatusPending,
		},
	}
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
	OrderStatusPaymentSent
	OrderStatusPaymentReversed
	OrderStatusPaymentReversedError
	OrderStatusPaymentError
	OrderStatusInPreparation
	OrderStatusReadyForDelivery
	OrderStatusSentForDelivery
	OrderStatusDelivered
	OrderStatusDeliveryError
	OrderStatusCanceled
	OrderStatusUnknown
)

func (os OrderStatus) ToString() string {
	return [...]string{
		"STARTED",
		"ADDING_ITEMS",
		"CONFIRMED",
		"PAID",
		"PAYMENT_SENT",
		"PAYMENT_REVERSED",
		"PAYMENT_REVERSED_ERROR",
		"PAYMENT_ERROR",
		"IN_PREPARATION",
		"READY_FOR_DELIVERY",
		"SENT_FOR_DELIVERY",
		"DELIVERED",
		"DELIVERY_ERROR",
		"CANCELED",
		"UNKNOWN",
	}[os]
}

func ToOrderStatus(status string) OrderStatus {
	switch status {
	case "STARTED":

		return OrderStatusStarted
	case "ADDING_ITEMS":
		return OrderStatusAddingItems
	case "CONFIRMED":

		return OrderStatusConfirmed
	case "IN_PREPARATION":

		return OrderStatusInPreparation
	case "READY_FOR_DELIVERY":

		return OrderStatusReadyForDelivery
	case "PAID":

		return OrderStatusPaid
	case "PAYMENT_SENT":
		return OrderStatusPaymentSent
	case "PAYMENT_REVERSED":

		return OrderStatusPaymentReversed
	case "PAYMENT_REVERSED_ERROR":

		return OrderStatusPaymentReversedError
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

		return OrderStatusUnknown
	}
}
