package dtos

import (
	"event-backend/entities"
)

type AuthRegisterRequestDto struct {
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (dto *AuthRegisterRequestDto) ToEntity() *entities.UserEntity {
	return &entities.UserEntity{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
