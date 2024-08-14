package ports

import "github.com/fabianogoes/fiap-challenge/domain/entities"

type PaymentClientPort interface {
	Pay(order *entities.Order, paymentMethod string) error
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

type PaymentPublisherPort interface {
	Publish(order *entities.Order, paymentMethod string) error
}

type PaymentReceiverPort interface {
	ReceiveCallback()
}
