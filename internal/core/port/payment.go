package port

import "github.com/fiap/challenge-gofood/internal/core/domain"

type PaymentClientPort interface {
	Pay(order *domain.Order) error
	Reverse(order *domain.Order) error
}

type PaymentUseCasePort interface {
	GetPaymentById(id uint) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) (*domain.Payment, error)
}

type PaymentRepositoryPort interface {
	GetPaymentById(id uint) (*domain.Payment, error)
	UpdatePayment(payment *domain.Payment) (*domain.Payment, error)
}
