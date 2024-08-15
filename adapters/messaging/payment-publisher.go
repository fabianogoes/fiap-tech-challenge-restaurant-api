package messaging

import (
	"encoding/json"
	"fmt"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository"
)

type PaymentPublisher struct {
	awsSQSClient     *AWSSQSClient
	outboxRepository *repository.OutboxRepository
}

func NewPaymentPublisher(awsSQSClient *AWSSQSClient, outboxRepository *repository.OutboxRepository) *PaymentPublisher {
	return &PaymentPublisher{awsSQSClient, outboxRepository}
}

func (p *PaymentPublisher) PublishPayment(order *entities.Order, paymentMethod string) error {
	fmt.Printf("Sending order %v to payment\n", order.ID)

	queueUrl := p.awsSQSClient.config.PaymentQueueUrl
	messageBody := toPaymentMessageBody(order, paymentMethod)
	if _, err := p.outboxRepository.CreateOutbox(order.ID, messageBody, queueUrl); err != nil {
		return fmt.Errorf("error creating outbox for order %v: %v", order.ID, err)
	}

	err := p.awsSQSClient.Publish(queueUrl, messageBody)
	if err != nil {
		return fmt.Errorf("error sending message to payment: %v", err)
	}

	if err := p.outboxRepository.DeleteOutbox(order.ID); err != nil {
		return fmt.Errorf("error deleting outbox for order %v: %v", order.ID, err)
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
