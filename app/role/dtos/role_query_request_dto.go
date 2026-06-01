package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type RoleQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
}

func AssignRoleQueryRequestDto(c *gin.Context) *RoleQueryRequestDto {
	return &RoleQueryRequestDto{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
