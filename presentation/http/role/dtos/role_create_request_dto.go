package dtos

import (
	"event-backend/entities"
)

type RoleCreateRequestDTO struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}

func (dto *RoleCreateRequestDTO) ToEntity() *entities.RoleEntity {
	return &entities.RoleEntity{
		Name: dto.Name,
	}
}
