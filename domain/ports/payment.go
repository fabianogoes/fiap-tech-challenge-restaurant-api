package ports

import "github.com/fabianogoes/fiap-challenge/domain/entities"

type PaymentUseCasePort interface {
	GetPaymentById(id uint) (*entities.Payment, error)
	UpdatePayment(payment *entities.Payment) (*entities.Payment, error)
}

type PaymentRepositoryPort interface {
	GetPaymentById(id uint) (*entities.Payment, error)
	UpdatePayment(payment *entities.Payment) (*entities.Payment, error)
}

type PaymentPublisherPort interface {
	PublishPayment(order *entities.Order, paymentMethod string) error
}

type PaymentReceiverPort interface {
	ReceivePaymentCallback()
}
