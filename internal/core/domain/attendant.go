package domain

import "time"

type Attendant struct {
	ID        int64
	Nome      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
