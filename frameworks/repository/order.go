package repository

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	dbo2 "github.com/fabianogoes/fiap-challenge/frameworks/repository/dbo"
	"time"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db             *gorm.DB
	itemRepository *OrderItemRepository
}

func NewOrderRepository(db *gorm.DB, itemRepo *OrderItemRepository) *OrderRepository {
	return &OrderRepository{db, itemRepo}
}

func (or *OrderRepository) CreateOrder(entity *entities.Order) (*entities.Order, error) {
	order := &dbo2.Order{
		CustomerID:  entity.Customer.ID,
		AttendantID: entity.Attendant.ID,
		Date:        time.Now(),
		Status:      entity.Status.ToString(),
		Payment:     dbo2.ToPaymentDBO(entity.Payment),
		Delivery:    dbo2.ToDeliveryDBO(entity.Delivery),
		Amount:      entity.Amount(),
	}

	if err := or.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order.ToEntity(), nil
}

func (or *OrderRepository) GetOrderById(id uint) (*entities.Order, error) {
	order := &dbo2.Order{}

	if err := or.db.Preload("Customer").
		Preload("Attendant").
		Preload("Payment").
		Preload("Delivery").
		Preload("Items").
		First(order, id).Error; err != nil {
		return nil, fmt.Errorf("error to find order with id %d - %v", id, err)
	}

	for _, item := range order.Items {
		product := &dbo2.Product{}
		if err := or.db.First(product, item.ProductID).Error; err != nil {
			return nil, fmt.Errorf("error to find product with id %d - %v", item.ProductID, err)
		}
		item.Product = product
	}

	return order.ToEntity(), nil
}

func (or *OrderRepository) UpdateOrder(order *entities.Order) (*entities.Order, error) {
	orderToUpdate := &dbo2.Order{}

	if err := or.db.Preload("Customer").Preload("Attendant").Preload("Payment").Preload("Items").
		First(orderToUpdate, order.ID).Error; err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		if item.ID == 0 {
			if err := or.itemRepository.CreateOrderItem(dbo2.ToOrderItemDBO(item)); err != nil {
				return nil, err
			}
		}
	}

	orderToUpdate.Amount = order.Amount()
	orderToUpdate.Status = order.Status.ToString()
	orderToUpdate.Payment = dbo2.ToPaymentDBO(order.Payment)

	if err := or.db.Save(orderToUpdate).Error; err != nil {
		return nil, err
	}

	return or.GetOrderById(order.ID)
}

func (or *OrderRepository) RemoveItemFromOrder(idItem uint) error {
	return or.itemRepository.Delete(idItem)
}

func (or *OrderRepository) GetOrderItemById(id uint) (*entities.OrderItem, error) {
	return or.itemRepository.GetOrderItemById(id)
}

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db}
}

func (oir *OrderItemRepository) CreateOrderItem(orderItem *dbo2.OrderItem) error {
	if err := oir.db.Create(orderItem).Error; err != nil {
		return err
	}

	return nil
}

func (oir *OrderItemRepository) Delete(idItem uint) error {
	if err := oir.db.Delete(&dbo2.OrderItem{}, idItem).Error; err != nil {
		return err
	}

	return nil
}

func (oir *OrderItemRepository) GetOrderItemById(id uint) (*entities.OrderItem, error) {
	orderItem := &dbo2.OrderItem{}

	if err := oir.db.Preload("Product").First(orderItem, id).Error; err != nil {
		return nil, fmt.Errorf("error to find order item with id %d - %v", id, err)
	}

	return orderItem.ToEntity(), nil
}
