package repository

import (
	"fmt"
	"time"

	"github.com/fiap/challenge-gofood/internal/core/domain"
	"gorm.io/gorm"
)

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
	Items       []*OrderItem
}

func (o *Order) ToModel() *domain.Order {
	var items []*domain.OrderItem
	var itemsTotal int

	for _, item := range o.Items {
		items = append(items, item.ToModel())
		itemsTotal += int(item.Quantity)
	}

	return &domain.Order{
		ID: o.ID,
		Customer: &domain.Customer{
			ID:   o.Customer.ID,
			Name: o.Customer.Name,
			CPF:  o.Customer.CPF,
		},
		Attendant: &domain.Attendant{
			ID:   o.Attendant.ID,
			Name: o.Attendant.Name,
		},
		Date:       o.Date,
		Status:     mapOrderStatus(o.Status),
		Payment:    o.Payment.ToModel(),
		Amount:     o.Amount,
		ItemsTotal: itemsTotal,
		Items:      items,
		CreatedAt:  o.CreatedAt,
		UpdatedAt:  o.UpdatedAt,
	}
}

func mapOrderStatus(status string) domain.OrderStatus {
	switch status {
	case "STARTED":
		return domain.OrderStatusStarted
	case "ADDING_ITEMS":
		return domain.OrderStatusAddingItems
	case "CONFIRMED":
		return domain.OrderStatusConfirmed
	case "PAID":
		return domain.OrderStatusPaid
	case "PAYMENT_REVERSED":
		return domain.OrderStatusPaymentReversed
	case "IN_PREPARATION":
		return domain.OrderStatusInPreparation
	case "READY_FOR_DELIVERY":
		return domain.OrderStatusReadyForDelivery
	case "SENT_FOR_DELIVERY":
		return domain.OrderStatusSentForDelivery
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
	Product   *Product
	Quantity  int64
	UnitPrice float64
}

func (i *OrderItem) ToModel() *domain.OrderItem {
	return &domain.OrderItem{
		ID:        i.ID,
		Product:   i.Product.ToModel(),
		Quantity:  int(i.Quantity),
		UnitPrice: i.UnitPrice,
	}
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
		CustomerID:  customerID,
		AttendantID: attendantID,
		Date:        time.Now(),
		Status:      orderStatus,
		Payment:     Payment{Status: paymentStatus},
		Amount:      0,
		Items:       []*OrderItem{},
	}

	if err := or.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order.ToModel(), nil
}

func (or *OrderRepository) GetOrderById(id uint) (*domain.Order, error) {
	order := &Order{}

	if err := or.db.Preload("Customer").Preload("Attendant").Preload("Payment").Preload("Items").
		First(order, id).Error; err != nil {
		return nil, fmt.Errorf("error to find order with id %d - %v", id, err)
	}

	for _, item := range order.Items {
		product := &Product{}
		if err := or.db.First(product, item.ProductID).Error; err != nil {
			return nil, fmt.Errorf("error to find product with id %d - %v", item.ProductID, err)
		}
		item.Product = product
	}

	return order.ToModel(), nil
}

func (or *OrderRepository) AddItemToOrder(order *domain.Order, product *domain.Product, quantity int) (*domain.Order, error) {
	orderItem := &OrderItem{
		OrderID:   order.ID,
		ProductID: product.ID,
		Quantity:  int64(quantity),
		UnitPrice: product.Price,
	}

	if err := or.db.Create(orderItem).Error; err != nil {
		return nil, err
	}

	return or.UpdateOrder(order)
}

func (or *OrderRepository) UpdateOrder(order *domain.Order) (*domain.Order, error) {
	orderToUpdate := &Order{}

	if err := or.db.Preload("Customer").Preload("Attendant").Preload("Payment").Preload("Items").
		First(orderToUpdate, order.ID).Error; err != nil {
		return nil, err
	}

	orderToUpdate.Amount = order.Amount
	orderToUpdate.Status = order.Status.ToString()
	orderToUpdate.Payment = mapPatmentEntity(order.Payment)

	fmt.Printf("Status: %s Payment: %s\n", orderToUpdate.Status, orderToUpdate.Payment.Status)
	fmt.Println(orderToUpdate.Payment)

	if err := or.db.Save(orderToUpdate).Error; err != nil {
		return nil, err
	}

	return or.GetOrderById(order.ID)
}
