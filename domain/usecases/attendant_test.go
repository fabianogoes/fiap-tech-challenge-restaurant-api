package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var attendantIDSuccess = uint(1)
var AttendantSuccess = &entities.Attendant{
	ID:        attendantIDSuccess,
	Name:      "Test Attendant",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestAttendantService_CreateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)

	repository.On("CreateAttendant", AttendantSuccess.Name).Return(AttendantSuccess, nil)
	service := NewAttendantService(repository)

	attendant, err := service.CreateAttendant(AttendantSuccess.Name)
	assert.NoError(t, err)
	assert.NotNil(t, attendant)
}

func TestAttendantService_GetAttendantById(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("GetAttendantById", AttendantSuccess.ID).Return(AttendantSuccess, nil)

	service := NewAttendantService(repository)

	attendant, err := service.GetAttendantById(AttendantSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, attendant)
}

func TestAttendantService_GetAttendants(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("GetAttendants").Return([]*entities.Attendant{AttendantSuccess}, nil)

	service := NewAttendantService(repository)

	attendants, err := service.GetAttendants()
	assert.NoError(t, err)
	assert.NotNil(t, attendants)
}

func TestAttendantService_UpdateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("UpdateAttendant", AttendantSuccess).Return(AttendantSuccess, nil)

	service := NewAttendantService(repository)

	updateAttendant, err := service.UpdateAttendant(AttendantSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, updateAttendant)
}
