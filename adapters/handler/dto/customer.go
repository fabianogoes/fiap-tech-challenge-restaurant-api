package dto

import "github.com/fiap/challenge-gofood/entities"

type GetCustomerResponse struct {
	ID        uint   `json:"id"`
	Nome      string `json:"name"`
	Email     string `json:"email"`
	CPF       string `json:"cpf"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToCustomerResponse(entity *entities.Customer) GetCustomerResponse {
	return GetCustomerResponse{
		ID:        entity.ID,
		Nome:      entity.Name,
		Email:     entity.Email,
		CPF:       entity.CPF,
		CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToCustomerResponses(customers []*entities.Customer) []GetCustomerResponse {
	var response []GetCustomerResponse
	for _, customer := range customers {
		response = append(response, ToCustomerResponse(customer))
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
