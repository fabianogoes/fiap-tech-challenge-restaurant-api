package service

import (
	"github.com/fiap/challenge-gofood/internal/domain/entity"
	"github.com/fiap/challenge-gofood/internal/domain/port"
)

type AttendantService struct {
	attendantRepository port.AttendantRepositoryPort
}

func NewAttendantService(cr port.AttendantRepositoryPort) *AttendantService {
	return &AttendantService{
		attendantRepository: cr,
	}
}

func (c *AttendantService) CreateAttendant(nome string) (*entity.Attendant, error) {
	return c.attendantRepository.CreateAttendant(nome)
}

func (c *AttendantService) GetAttendantById(id uint) (*entity.Attendant, error) {
	return c.attendantRepository.GetAttendantById(id)
}

func (c *AttendantService) GetAttendants() ([]*entity.Attendant, error) {
	return c.attendantRepository.GetAttendants()
}

func (c *AttendantService) UpdateAttendant(attendant *entity.Attendant) (*entity.Attendant, error) {
	return c.attendantRepository.UpdateAttendant(attendant)
}

func (c *AttendantService) DeleteAttendant(id uint) error {
	return c.attendantRepository.DeleteAttendant(id)
}
