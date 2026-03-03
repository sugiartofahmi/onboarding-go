package dtos

import (
	"github.com/google/uuid"

	"event-backend/entities"
)

type UserCreateRequestDTO struct {
	Name     string    `json:"name" binding:"required,min=3,max=255"`
	Email    string    `json:"email" binding:"required,email,max=255"`
	Password string    `json:"password" binding:"required,min=6,max=255"`
	RoleId   uuid.UUID `json:"role_id" binding:"required"`
}

func (dto *UserCreateRequestDTO) ToEntity() *entities.UserEntity {
	return &entities.UserEntity{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		RoleId:   dto.RoleId,
	}
}
