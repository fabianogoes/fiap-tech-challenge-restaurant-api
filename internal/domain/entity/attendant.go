package entity

import "time"

type Attendant struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAttendant(name string) (*Attendant, error) {
	return &Attendant{
		Name: name,
	}, nil
}
