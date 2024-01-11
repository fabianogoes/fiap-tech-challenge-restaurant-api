package dto

import "github.com/fiap/challenge-gofood/internal/domain/entity"

type GetProductResponse struct {
	ID         uint              `json:"id"`
	Name       string            `json:"name"`
	Price      float64           `json:"price"`
	Category   *CategoryResponse `json:"category"`
	CreadtedAt string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
}

func ToProductResponse(entity *entity.Product) GetProductResponse {
	return GetProductResponse{
		ID:         entity.ID,
		Name:       entity.Name,
		Price:      entity.Price,
		Category:   CategoryToResponse(entity.Category),
		CreadtedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToProductResponses(entities []*entity.Product) []GetProductResponse {
	var response []GetProductResponse
	for _, entity := range entities {
		response = append(response, ToProductResponse(entity))
	}
	return response
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CategoryToResponse(category *entity.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"categoryID"`
}

type UpdateProductRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}
