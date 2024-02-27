package usecases

import (
	"github.com/fiap/challenge-gofood/entities"
	"github.com/fiap/challenge-gofood/interfaces"
)

type AttendantService struct {
	attendantRepository interfaces.AttendantRepositoryPort
}

func NewAttendantService(cr interfaces.AttendantRepositoryPort) *AttendantService {
	return &AttendantService{
		attendantRepository: cr,
	}
}

func (c *AttendantService) CreateAttendant(nome string) (*entities.Attendant, error) {
	return c.attendantRepository.CreateAttendant(nome)
}

func (c *AttendantService) GetAttendantById(id uint) (*entities.Attendant, error) {
	return c.attendantRepository.GetAttendantById(id)
}

func (c *AttendantService) GetAttendants() ([]*entities.Attendant, error) {
	return c.attendantRepository.GetAttendants()
}

func (c *AttendantService) UpdateAttendant(attendant *entities.Attendant) (*entities.Attendant, error) {
	return c.attendantRepository.UpdateAttendant(attendant)
}

func (c *AttendantService) DeleteAttendant(id uint) error {
	return c.attendantRepository.DeleteAttendant(id)
}
