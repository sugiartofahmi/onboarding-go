package dtos

import (
	"time"

	"event-backend/entities"

	"github.com/google/uuid"
)

type CategoryDetailResponseDto struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CategoryDetailResponseDtoFromEntity(entity entities.CategoryEntity) CategoryDetailResponseDto {
	return CategoryDetailResponseDto{
		Id:        entity.Id,
		Name:      entity.Name,
		Slug:      entity.Slug,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
