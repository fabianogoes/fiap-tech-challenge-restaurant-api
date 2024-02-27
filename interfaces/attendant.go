package interfaces

import "github.com/fabianogoes/fiap-challenge/entities"

// Primary ports to Attendant

type AttendantUseCasePort interface {
	CreateAttendant(nome string) (*entities.Attendant, error)
	GetAttendantById(id uint) (*entities.Attendant, error)
	GetAttendants() ([]*entities.Attendant, error)
	UpdateAttendant(attendant *entities.Attendant) (*entities.Attendant, error)
	DeleteAttendant(id uint) error
}

// Secondary ports to Attendant

type AttendantRepositoryPort interface {
	CreateAttendant(nome string) (*entities.Attendant, error)
	GetAttendantById(id uint) (*entities.Attendant, error)
	GetAttendantByName(name string) (*entities.Attendant, error)
	GetAttendants() ([]*entities.Attendant, error)
	UpdateAttendant(attendant *entities.Attendant) (*entities.Attendant, error)
	DeleteAttendant(id uint) error
}
