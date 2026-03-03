package dtos

import (
	"event-backend/entities"
	"time"

	"github.com/google/uuid"
)

type UserDetailResponseDto struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func UserDetailResponseDtoFromEntity(entity *entities.UserEntity) *UserDetailResponseDto {
	if entity == nil {
		return nil
	}
	return &UserDetailResponseDto{
		Id:        entity.Id,
		Name:      entity.Name,
		Email:     entity.Email,
		Role:      entity.Role.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}