package messaging

import (
	"encoding/json"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
)

type PaymentPublisher struct {
	awsSQSClient *AWSSQSClient
}

func NewPaymentPublisher(awsSQSClient *AWSSQSClient) *PaymentPublisher {
	return &PaymentPublisher{awsSQSClient}
}

func (m *PaymentPublisher) PublishPayment(order *entities.Order, paymentMethod string) error {
	fmt.Printf("Sending order %v to payment\n", order.ID)

	queueName := m.awsSQSClient.config.PaymentQueueUrl
	err := m.awsSQSClient.Publish(queueName, toPaymentMessageBody(order, paymentMethod))
	if err != nil {
		return fmt.Errorf("error sending message to payment: %v", err)
	}

	return nil
}

func toPaymentMessageBody(order *entities.Order, paymentMethod string) string {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"orderId": order.ID,
		"method":  paymentMethod,
		"value":   order.Amount(),
		"date":    order.Date.Format("2006-01-02T15:04:05"),
	})

	return string(jsonBytes)
}
