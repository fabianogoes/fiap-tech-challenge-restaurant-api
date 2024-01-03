package port

import "github.com/fiap/challenge-gofood/internal/core/domain"

// Primary ports to Attendant

type AttendantUseCasePort interface {
	CreateAttendant(nome string) (*domain.Attendant, error)
	GetAttendantById(id uint) (*domain.Attendant, error)
	GetAttendants() ([]*domain.Attendant, error)
	UpdateAttendant(attendant *domain.Attendant) (*domain.Attendant, error)
	DeleteAttendant(id uint) error
}

// Secondary ports to Attendant

type AttendantRepositoryPort interface {
	CreateAttendant(nome string) (*domain.Attendant, error)
	GetAttendantById(id uint) (*domain.Attendant, error)
	GetAttendants() ([]*domain.Attendant, error)
	UpdateAttendant(attendant *domain.Attendant) (*domain.Attendant, error)
	DeleteAttendant(id uint) error
}
