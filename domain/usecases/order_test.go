package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"time"
)

var OrderIDSuccess = uint(1)
var OrderStarted = &entities.Order{
	ID:        OrderIDSuccess,
	Customer:  CustomerSuccess,
	Attendant: AttendantSuccess,
	Date:      time.Now(),
	Status:    entities.OrderStatusStarted,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}
