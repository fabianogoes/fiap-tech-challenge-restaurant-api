package entities

import "time"

type Outbox struct {
	ID          uint
	CreatedAt   time.Time
	MessageBody string
	QueueUrl    string
}
