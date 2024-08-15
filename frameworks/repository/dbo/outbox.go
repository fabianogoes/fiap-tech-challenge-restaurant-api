package dbo

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"time"
)

type Outbox struct {
	ID          uint
	MessageBody string
	CreatedAt   time.Time
	QueueUrl    string
}

func (o *Outbox) ToEntity() *entities.Outbox {
	return &entities.Outbox{
		ID:          o.ID,
		CreatedAt:   o.CreatedAt,
		MessageBody: o.MessageBody,
		QueueUrl:    o.QueueUrl,
	}
}

func ToOutboxDBO(orderID uint, messageBody string, queueUrl string) Outbox {
	return Outbox{
		ID:          orderID,
		MessageBody: messageBody,
		CreatedAt:   time.Now(),
		QueueUrl:    queueUrl,
	}
}
