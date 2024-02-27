package dto

import "github.com/fiap/challenge-gofood/entities"

type GetProductResponse struct {
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	Price     float64           `json:"price"`
	Category  *CategoryResponse `json:"category"`
	CreatedAt string            `json:"createdAt"`
	UpdatedAt string            `json:"updatedAt"`
}

func ToProductResponse(entity *entities.Product) GetProductResponse {
	return GetProductResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Price:     entity.Price,
		Category:  CategoryToResponse(entity.Category),
		CreatedAt: entity.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: entity.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToProductResponses(products []*entities.Product) []GetProductResponse {
	var response []GetProductResponse
	for _, product := range products {
		response = append(response, ToProductResponse(product))
	}
	return response
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CategoryToResponse(category *entities.Category) *CategoryResponse {
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
