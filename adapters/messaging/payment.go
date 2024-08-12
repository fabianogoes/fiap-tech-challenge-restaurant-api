package messaging

import (
	"encoding/json"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
)

type PaymentMessaging struct {
}

func NewPaymentMessaging() *PaymentMessaging {
	return &PaymentMessaging{}
}

func (m *PaymentMessaging) Send(order *entities.Order, paymentMethod string) error {
	fmt.Printf("Sending order %+v\n", order)

	queueName := "order-payment-queue"
	err := Publish(queueName, toMessageBody(order, paymentMethod))
	if err != nil {
		return fmt.Errorf("error sending message to queue: %v", err)
	}

	return nil
}

func toMessageBody(order *entities.Order, paymentMethod string) string {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"orderId": order.ID,
		"method":  paymentMethod,
		"value":   order.Amount(),
		"date":    order.Date.Format("2006-01-02T15:04:05"),
	})

	return string(jsonBytes)
}
