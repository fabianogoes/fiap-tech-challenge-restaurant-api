package usecases

import (
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAttendantService_CreateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)

	repository.On("CreateAttendant", domain.AttendantSuccess.Name).Return(domain.AttendantSuccess, nil)
	service := NewAttendantService(repository)

	attendant, err := service.CreateAttendant(domain.AttendantSuccess.Name)
	assert.NoError(t, err)
	assert.NotNil(t, attendant)
}

func TestAttendantService_GetAttendantById(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("GetAttendantById", domain.AttendantSuccess.ID).Return(domain.AttendantSuccess, nil)

	service := NewAttendantService(repository)

	attendant, err := service.GetAttendantById(domain.AttendantSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, attendant)
}

func TestAttendantService_GetAttendants(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("GetAttendants").Return([]*entities.Attendant{domain.AttendantSuccess}, nil)

	service := NewAttendantService(repository)

	attendants, err := service.GetAttendants()
	assert.NoError(t, err)
	assert.NotNil(t, attendants)
}

func TestAttendantService_UpdateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("UpdateAttendant", domain.AttendantSuccess).Return(domain.AttendantSuccess, nil)

	service := NewAttendantService(repository)

	updateAttendant, err := service.UpdateAttendant(domain.AttendantSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, updateAttendant)
}

func TestAttendantService_DeleteAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	repository.On("DeleteAttendant", mock.Anything).Return(nil)

	service := NewAttendantService(repository)

	err := service.DeleteAttendant(1)
	assert.NoError(t, err)
}
