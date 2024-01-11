package dto

import "github.com/fiap/challenge-gofood/internal/domain/entity"

type GetCustomerResponse struct {
	ID        uint   `json:"id"`
	Nome      string `json:"name"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToCustomerResponse(entity *entity.Customer) GetCustomerResponse {
	return GetCustomerResponse{
		ID:        entity.ID,
		Nome:      entity.Name,
		Email:     entity.Email,
		CPF:       entity.CPF,
		CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCustomerResponses(entities []*entity.Customer) []GetCustomerResponse {
	var response []GetCustomerResponse
	for _, entity := range entities {
		response = append(response, ToCustomerResponse(entity))
	}
	return response
}

type CreateCustomerRequest struct {
	Nome  string `json:"name"`
	Email string `json:"email"`
	CPF   string `json:"cpf"`
}

type UpdateCustomerRequest struct {
	Nome  string `json:"name"`
	Email string `json:"email"`
}
