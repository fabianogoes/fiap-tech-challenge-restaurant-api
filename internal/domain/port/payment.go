package port

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

type PaymentClientPort interface {
	Pay(order *entity.Order) error
	Reverse(order *entity.Order) error
}

type PaymentUseCasePort interface {
	GetPaymentById(id uint) (*entity.Payment, error)
	UpdatePayment(payment *entity.Payment) (*entity.Payment, error)
}

type PaymentRepositoryPort interface {
	GetPaymentById(id uint) (*entity.Payment, error)
	UpdatePayment(payment *entity.Payment) (*entity.Payment, error)
}
