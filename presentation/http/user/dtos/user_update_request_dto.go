package dtos

import (
	"github.com/google/uuid"

	"event-backend/entities"
)

type UserUpdateRequestDTO struct {
	Id       uuid.UUID  `json:"id" binding:"required"`
	Name     string     `json:"name" binding:"required,min=3,max=255"`
	Email    string     `json:"email" binding:"required,email,max=255"`
	RoleId   *uuid.UUID `json:"role_id"`
}

func (dto *UserUpdateRequestDTO) ToEntity(existing *entities.UserEntity) *entities.UserEntity {
	if dto.Name != "" {
		existing.Name = dto.Name
	}
	if dto.Email != "" {
		existing.Email = dto.Email
	}
	if dto.RoleId != nil {
		existing.RoleId = *dto.RoleId
	}
	return existing
}
