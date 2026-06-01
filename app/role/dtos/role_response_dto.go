package dtos

import (
	"time"

	"event-backend/entities"

	"github.com/google/uuid"
)

type RoleResponseDto struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func RoleResponseDtoFromEntities(entities []*entities.RoleEntity) []RoleResponseDto {
	roleResponses := make([]RoleResponseDto, len(entities))
	for i, role := range entities {
		roleResponses[i] = RoleResponseDto{
			Id:        role.Id,
			Name:      role.Name,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
		}
	}

	return roleResponses
}

func RoleDetailResponseDtoFromEntity(entity entities.RoleEntity) RoleResponseDto {
	return RoleResponseDto{
		Id:        entity.Id,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
