package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

func TestCustomer_GetCustomers(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomers").Return([]*entities.Customer{domain.CustomerSuccess}, nil)

	setup := SetupTest()
	setup.GET("/customers", handler.GetCustomers)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_GetCustomersInternalServerError(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomers").Return([]*entities.Customer{domain.CustomerSuccess}, errors.New("not found"))

	setup := SetupTest()
	setup.GET("/customers", handler.GetCustomers)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_GetCustomersNoContent(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomers").Return([]*entities.Customer{}, nil)

	setup := SetupTest()
	setup.GET("/customers", handler.GetCustomers)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_GetCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	setup := SetupTest()
	setup.GET("/customers/:id", handler.GetCustomer)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers/1"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_GetCustomerBadRequest(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)

	setup := SetupTest()
	setup.GET("/customers/:id", handler.GetCustomer)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers/x"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_GetCustomerInternalServerError(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, errors.New("not found"))

	setup := SetupTest()
	setup.GET("/customers/:id", handler.GetCustomer)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers/1"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_GetCustomerByCPF(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerByCPF", mock.Anything).Return(domain.CustomerSuccess, nil)

	setup := SetupTest()
	setup.GET("/customers/cpf/:cpf", handler.GetCustomerByCPF)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers/cpf/123"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_GetCustomerByCPFInternalServerError(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerByCPF", mock.Anything).Return(domain.CustomerSuccess, errors.New("not found"))

	setup := SetupTest()
	setup.GET("/customers/cpf/:cpf", handler.GetCustomerByCPF)
	request, err := http.NewRequest("GET", fmt.Sprintf("/customers/cpf/123"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_CreateCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("CreateCustomer", mock.Anything, mock.Anything, mock.Anything).
		Return(domain.CustomerSuccess, nil)

	payload := dto.CreateCustomerRequest{Name: "test", CPF: "123", Email: "test@test.com"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/customers/", handler.CreateCustomer)
	request, err := http.NewRequest("POST", fmt.Sprintf("/customers/"), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusCreated, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_CreateCustomerBadRequest(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("CreateCustomer", mock.Anything, mock.Anything, mock.Anything).
		Return(domain.CustomerSuccess, nil)

	setup := SetupTest()
	setup.POST("/customers/", handler.CreateCustomer)
	request, err := http.NewRequest("POST", fmt.Sprintf("/customers/"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_CreateCustomerInternalServerError(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("CreateCustomer", mock.Anything, mock.Anything, mock.Anything).
		Return(domain.CustomerSuccess, errors.New("not found"))

	payload := dto.CreateCustomerRequest{Name: "test", CPF: "123", Email: "test@test.com"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/customers/", handler.CreateCustomer)
	request, err := http.NewRequest("POST", fmt.Sprintf("/customers/"), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_UpdateCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)
	repository.On("UpdateCustomer", mock.Anything).Return(domain.CustomerSuccess, nil)

	payload := dto.CreateCustomerRequest{Name: "test", CPF: "123", Email: "test@test.com"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.PUT("/customers/:id", handler.UpdateCustomer)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/customers/1"), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusAccepted, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_UpdateCustomerInternalServerError(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, errors.New("not found"))
	repository.On("UpdateCustomer", mock.Anything).Return(domain.CustomerSuccess, nil)

	payload := dto.CreateCustomerRequest{Name: "test", CPF: "123", Email: "test@test.com"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.PUT("/customers/:id", handler.UpdateCustomer)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/customers/1"), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_UpdateCustomerBadRequestId(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, errors.New("not found"))
	repository.On("UpdateCustomer", mock.Anything).Return(domain.CustomerSuccess, nil)

	payload := dto.CreateCustomerRequest{Name: "test", CPF: "123", Email: "test@test.com"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.PUT("/customers/:id", handler.UpdateCustomer)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/customers/x"), readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_UpdateCustomerBadRequestJson(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerById", mock.Anything).Return(domain.CustomerSuccess, nil)
	repository.On("UpdateCustomer", mock.Anything).Return(domain.CustomerSuccess, nil)

	setup := SetupTest()
	setup.PUT("/customers/:id", handler.UpdateCustomer)
	request, err := http.NewRequest("PUT", fmt.Sprintf("/customers/x"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_DeleteCustomer(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("DeleteCustomer", mock.Anything).Return(nil)

	setup := SetupTest()
	setup.DELETE("/customers/:id", handler.DeleteCustomer)
	request, err := http.NewRequest("DELETE", fmt.Sprintf("/customers/1"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusNoContent, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_DeleteCustomerInternalServerError(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("DeleteCustomer", mock.Anything).Return(errors.New("error"))

	setup := SetupTest()
	setup.DELETE("/customers/:id", handler.DeleteCustomer)
	request, err := http.NewRequest("DELETE", fmt.Sprintf("/customers/1"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusInternalServerError, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_DeleteCustomerBadRequest(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("DeleteCustomer", mock.Anything).Return(errors.New("error"))

	setup := SetupTest()
	setup.DELETE("/customers/:id", handler.DeleteCustomer)
	request, err := http.NewRequest("DELETE", fmt.Sprintf("/customers/x"), nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_SignIn(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerByCPF", mock.Anything).Return(domain.CustomerSuccess, nil)

	payload := dto.TokenRequest{CPF: "123"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/customers/sign-in", handler.SignIn)
	request, err := http.NewRequest("POST", "/customers/sign-in", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_SignInUnprocessableEntity(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerByCPF", mock.Anything).Return(domain.CustomerSuccess, errors.New("not found"))

	payload := dto.TokenRequest{CPF: "123"}
	jsonRequest, _ := json.Marshal(payload)
	readerPayload := bytes.NewReader(jsonRequest)

	setup := SetupTest()
	setup.POST("/customers/sign-in", handler.SignIn)
	request, err := http.NewRequest("POST", "/customers/sign-in", readerPayload)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusUnprocessableEntity, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestCustomer_SignInBadRequest(t *testing.T) {
	repository := new(domain.CustomerRepositoryMock)
	useCase := usecases.NewCustomerService(repository)
	config, _ := entities.NewConfig()
	handler := NewCustomerHandler(useCase, config)

	repository.On("GetCustomerByCPF", mock.Anything).Return(domain.CustomerSuccess, errors.New("not found"))

	setup := SetupTest()
	setup.POST("/customers/sign-in", handler.SignIn)
	request, err := http.NewRequest("POST", "/customers/sign-in", nil)
	assert.NoError(t, err)

	response := httptest.NewRecorder()
	setup.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}
