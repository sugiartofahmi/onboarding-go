package dtos

import (
	"github.com/google/uuid"

	"event-backend/entities"
	"event-backend/infrastructure/utils"
)

type CategoryUpdateRequestDTO struct {
	CategoryCreateRequestDTO
	Id uuid.UUID `json:"-"`
}


func (dto *CategoryUpdateRequestDTO) ToEntity(existingEntity *entities.CategoryEntity) *entities.CategoryEntity {
	existingEntity.Name = dto.Name
	existingEntity.Slug = utils.GenerateSlug(dto.Name)
	return existingEntity
}