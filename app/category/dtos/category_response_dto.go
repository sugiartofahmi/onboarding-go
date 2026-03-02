package dtos

import (
	"time"

	"event-backend/entities"

	"github.com/google/uuid"
)

type CategoryResponseDto struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CategoryResponseDtoFromEntities(entities []*entities.CategoryEntity) []CategoryResponseDto {
	categoryResponses := make([]CategoryResponseDto, len(entities))
	for i, category := range entities {
		categoryResponses[i] = CategoryResponseDto{
			Id:        category.Id,
			Name:      category.Name,
			Slug:      category.Slug,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
	}

	return categoryResponses
}
