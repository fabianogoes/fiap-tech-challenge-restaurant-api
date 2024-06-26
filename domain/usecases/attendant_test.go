package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var attendantIDSuccess = uint(1)
var attendant = &entities.Attendant{
	ID:        attendantIDSuccess,
	Name:      "Test Attendant",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestAttendantService_CreateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)

	repository.On("CreateAttendant", attendant.Name).Return(attendant, nil)
	service := NewAttendantService(repository)

	attendant, err := service.CreateAttendant(attendant.Name)
	assert.Nil(t, err)
	assert.NotNil(t, attendant)
}

func TestAttendantService_GetAttendantById(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("GetAttendantById", attendant.ID).Return(attendant, nil)

	service := NewAttendantService(repository)

	attendant, err := service.GetAttendantById(attendant.ID)
	assert.NoError(t, err)
	assert.NotNil(t, attendant)
}

func TestAttendantService_GetAttendants(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("GetAttendants").Return([]*entities.Attendant{attendant}, nil)

	service := NewAttendantService(repository)

	attendants, err := service.GetAttendants()
	assert.NoError(t, err)
	assert.NotNil(t, attendants)
}

func TestAttendantService_UpdateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("UpdateAttendant", attendant).Return(attendant, nil)

	service := NewAttendantService(repository)

	updateAttendant, err := service.UpdateAttendant(attendant)
	assert.NoError(t, err)
	assert.NotNil(t, updateAttendant)
}
