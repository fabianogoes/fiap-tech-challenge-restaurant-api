package domain

import "time"

type Customer struct {
	ID        int64
	Nome      string
	Email     string
	CPF       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
