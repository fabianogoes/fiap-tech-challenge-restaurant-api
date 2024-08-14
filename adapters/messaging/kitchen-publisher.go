package messaging

import (
	"encoding/json"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
)

type KitchenPublisher struct {
	awsSQSClient *AWSSQSClient
}

func NewKitchenPublisher(awsSQSClient *AWSSQSClient) *KitchenPublisher {
	return &KitchenPublisher{awsSQSClient}
}

func (m *KitchenPublisher) PublishKitchen(order *entities.Order) error {
	fmt.Printf("Sending order %v to kitchen\n", order.ID)

	queueName := m.awsSQSClient.config.KitchenQueueUrl
	err := m.awsSQSClient.Publish(queueName, toKitchenMessageBody(order))
	if err != nil {
		return fmt.Errorf("error sending message to kitchen: %v", err)
	}

	return nil
}

func toKitchenMessageBody(order *entities.Order) string {
	kitchenDTO := toOrderKitchenDTO(order)
	jsonBytes, _ := json.Marshal(kitchenDTO)
	return string(jsonBytes)
}

type OrderKitchenDTO struct {
	ID    uint                `json:"id"`
	Items []*OrderKitchenItem `json:"items"`
}

type OrderKitchenItem struct {
	ID       uint            `json:"id"`
	Product  *ProductKitchen `json:"product"`
	Quantity int             `json:"quantity"`
}

type ProductKitchen struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func toOrderKitchenDTO(order *entities.Order) *OrderKitchenDTO {
	items := make([]*OrderKitchenItem, 0)
	for _, item := range order.Items {
		items = append(items, &OrderKitchenItem{
			ID: item.ID,
			Product: &ProductKitchen{
				ID:       item.Product.ID,
				Name:     item.Product.Name,
				Category: item.Product.Category.Name,
			},
			Quantity: item.Quantity,
		})
	}

	return &OrderKitchenDTO{
		ID:    order.ID,
		Items: items,
	}
}
