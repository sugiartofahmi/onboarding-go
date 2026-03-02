package dtos

import (
	"event-backend/entities"
	"event-backend/infrastructure/utils"
)

type CategoryCreateRequestDTO struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}

func (dto *CategoryCreateRequestDTO) ToEntity() *entities.CategoryEntity {
	return &entities.CategoryEntity{
		Name: dto.Name,
		Slug: utils.GenerateSlug(dto.Name),
	}
}