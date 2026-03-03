package dtos

import (
	"time"

	"event-backend/entities"

	"github.com/google/uuid"
)

type UserResponseDto struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserResponseDtoFromEntities(entities []*entities.UserEntity) []UserResponseDto {
	userResponses := make([]UserResponseDto, len(entities))
	for i, user := range entities {
		userResponses[i] = UserResponseDto{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}
	return userResponses
}
