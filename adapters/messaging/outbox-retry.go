package messaging

import (
	"fmt"
	"github.com/fabianogoes/fiap-challenge/frameworks/repository"
	"log/slog"
)

type OutboxRetry struct {
	sqsClient        *AWSSQSClient
	outboxRepository *repository.OutboxRepository
}

func NewOutboxRetry(sqsClient *AWSSQSClient, outboxRepository *repository.OutboxRepository) *OutboxRetry {
	return &OutboxRetry{sqsClient, outboxRepository}
}

func (o *OutboxRetry) Retry() {
	slog.Info("running outbox retry...")
	list, err := o.outboxRepository.GetAll()
	if err != nil {
		slog.Error(fmt.Sprintf("error retrieving all outbox list - %v", err))
		return
	}

	for _, outbox := range list {
		err := o.sqsClient.Publish(outbox.QueueUrl, outbox.MessageBody)
		if err != nil {
			slog.Error(fmt.Sprintf("error publishing outbox to quene %v, message - %v", outbox.QueueUrl, err))
			return
		}

		if err := o.outboxRepository.DeleteOutbox(outbox.ID); err != nil {
			slog.Error(fmt.Sprintf("error deleting outbox ID %v - %v \n", outbox.ID, err))
		}
	}
}
