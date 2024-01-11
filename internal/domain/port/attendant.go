package port

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
)

// Primary ports to Attendant

type AttendantUseCasePort interface {
	CreateAttendant(nome string) (*entity.Attendant, error)
	GetAttendantById(id uint) (*entity.Attendant, error)
	GetAttendants() ([]*entity.Attendant, error)
	UpdateAttendant(attendant *entity.Attendant) (*entity.Attendant, error)
	DeleteAttendant(id uint) error
}

// Secondary ports to Attendant

type AttendantRepositoryPort interface {
	CreateAttendant(nome string) (*entity.Attendant, error)
	GetAttendantById(id uint) (*entity.Attendant, error)
	GetAttendantByName(name string) (*entity.Attendant, error)
	GetAttendants() ([]*entity.Attendant, error)
	UpdateAttendant(attendant *entity.Attendant) (*entity.Attendant, error)
	DeleteAttendant(id uint) error
}
