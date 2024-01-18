package dto

import "github.com/fiap/challenge-gofood/internal/core/domain"

type CreateAttendantRequest struct {
	Name string `json:"name"`
}

type UpdateAttendantRequest struct {
	Nome string `json:"name"`
}

type GetAttendantResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ToAttendantResponse(entity *domain.Attendant) GetAttendantResponse {
	return GetAttendantResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToAttendantResponses(entities []*domain.Attendant) []GetAttendantResponse {
	var response []GetAttendantResponse
	for _, entity := range entities {
		response = append(response, ToAttendantResponse(entity))
	}
	return response
}
