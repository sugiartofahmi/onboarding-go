package dtos

import (
	"github.com/google/uuid"

	"event-backend/entities"
)

type RoleUpdateRequestDTO struct {
	RoleCreateRequestDTO
	Id uuid.UUID `json:"-"`
}

func (dto *RoleUpdateRequestDTO) ToEntity(existingEntity *entities.RoleEntity) *entities.RoleEntity {
	existingEntity.Name = dto.Name
	return existingEntity
}
