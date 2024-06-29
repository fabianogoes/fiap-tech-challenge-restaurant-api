package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/fabianogoes/fiap-challenge/domain"
	"github.com/fabianogoes/fiap-challenge/domain/entities"
	"github.com/fabianogoes/fiap-challenge/domain/usecases"
	"github.com/fabianogoes/fiap-challenge/frameworks/rest/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAttendantHandler_GetAttendants(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("GetAttendants").Return([]*entities.Attendant{domain.AttendantSuccess}, nil)

	setup := SetupTest()
	setup.GET("/attendants", handler.GetAttendants)
	request, err := http.NewRequest("GET", "/attendants", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_GetAttendantsInternalServerError(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("GetAttendants").Return(nil, errors.New("empty"))

	setup := SetupTest()
	setup.GET("/attendants", handler.GetAttendants)
	request, err := http.NewRequest("GET", "/attendants", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_GetAttendantsStatusNoContent(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("GetAttendants").Return([]*entities.Attendant{}, nil)

	setup := SetupTest()
	setup.GET("/attendants", handler.GetAttendants)
	request, err := http.NewRequest("GET", "/attendants", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_GetAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)

	setup := SetupTest()
	setup.GET("/attendants/:id", handler.GetAttendant)
	request, err := http.NewRequest("GET", "/attendants/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_GetAttendantInternalServerError(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("GetAttendantById", mock.Anything).Return(nil, errors.New("not found"))

	setup := SetupTest()
	setup.GET("/attendants/:id", handler.GetAttendant)
	request, err := http.NewRequest("GET", "/attendants/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_GetAttendantBadRequest(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	setup := SetupTest()
	setup.GET("/attendants/:id", handler.GetAttendant)
	request, err := http.NewRequest("GET", "/attendants/xxx", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_CreateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	payload := dto.CreateAttendantRequest{Name: "test"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	repository.On("CreateAttendant", mock.Anything).Return(domain.AttendantSuccess, nil)

	setup := SetupTest()
	setup.POST("/attendants/", handler.CreateAttendant)
	request, err := http.NewRequest("POST", "/attendants/", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_CreateAttendantBadRequest(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	setup := SetupTest()
	setup.POST("/attendants/", handler.CreateAttendant)
	request, err := http.NewRequest("POST", "/attendants/", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_CreateAttendantInternalServerError(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("CreateAttendant", mock.Anything).Return(nil, errors.New("error"))

	payload := dto.CreateAttendantRequest{Name: "test"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/attendants/", handler.CreateAttendant)
	request, err := http.NewRequest("POST", "/attendants/", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_UpdateAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	payload := dto.UpdateAttendantRequest{Name: "test"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	repository.On("GetAttendantById", mock.Anything).Return(domain.AttendantSuccess, nil)
	repository.On("UpdateAttendant", mock.Anything).Return(domain.AttendantSuccess, nil)

	setup := SetupTest()
	setup.PUT("/attendants/:id", handler.UpdateAttendant)
	request, err := http.NewRequest("PUT", "/attendants/1", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_UpdateAttendantBadRequest(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	setup := SetupTest()
	setup.PUT("/attendants/:id", handler.UpdateAttendant)
	request, err := http.NewRequest("PUT", "/attendants/xxx", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_UpdateAttendantBadRequestJson(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	setup := SetupTest()
	setup.PUT("/attendants/:id", handler.UpdateAttendant)
	request, err := http.NewRequest("PUT", "/attendants/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_UpdateAttendantInternalServerError(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("GetAttendantById", mock.Anything).Return(nil, errors.New("not found"))

	setup := SetupTest()
	setup.PUT("/attendants/:id", handler.UpdateAttendant)
	request, err := http.NewRequest("PUT", "/attendants/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_DeleteAttendantBadRequest(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	setup := SetupTest()
	setup.PUT("/attendants/:id", handler.DeleteAttendant)
	request, err := http.NewRequest("PUT", "/attendants/xxx", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_DeleteAttendant(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("DeleteAttendant", mock.Anything).Return(nil)

	setup := SetupTest()
	setup.DELETE("/attendants/:id", handler.DeleteAttendant)
	request, err := http.NewRequest("DELETE", "/attendants/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestAttendantHandler_DeleteAttendantInternalServerError(t *testing.T) {
	repository := new(domain.AttendantRepositoryMock)
	useCase := usecases.NewAttendantService(repository)
	handler := NewAttendantHandler(useCase)

	repository.On("DeleteAttendant", mock.Anything).Return(errors.New("not found"))

	setup := SetupTest()
	setup.DELETE("/attendants/:id", handler.DeleteAttendant)
	request, err := http.NewRequest("DELETE", "/attendants/1", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}
