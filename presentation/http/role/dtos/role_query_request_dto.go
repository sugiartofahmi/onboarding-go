package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type RoleQueryRequestDTO struct {
	infradtos.PaginationQueryRequestDto
}

func AssignRoleQueryRequestDTO(c *gin.Context) *RoleQueryRequestDTO {
	return &RoleQueryRequestDTO{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
