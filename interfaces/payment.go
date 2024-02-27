package interfaces

import "github.com/fabianogoes/fiap-challenge/entities"

type PaymentClientPort interface {
	Pay(order *entities.Order) error
	Reverse(order *entities.Order) error
}

type PaymentUseCasePort interface {
	GetPaymentById(id uint) (*entities.Payment, error)
	UpdatePayment(payment *entities.Payment) (*entities.Payment, error)
}

type PaymentRepositoryPort interface {
	GetPaymentById(id uint) (*entities.Payment, error)
	UpdatePayment(payment *entities.Payment) (*entities.Payment, error)
}
