package dtos

import (
	"github.com/google/uuid"

	"event-backend/entities"
)

type RoleUpdateRequestDto struct {
	RoleCreateRequestDto
	Id uuid.UUID `json:"-"`
}

func (dto *RoleUpdateRequestDto) ToEntity(existingEntity *entities.RoleEntity) *entities.RoleEntity {
	existingEntity.Name = dto.Name
	return existingEntity
}
