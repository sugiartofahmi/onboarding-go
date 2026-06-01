package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	infradtos "event-backend/infrastructure/dtos"
)

type UserQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
	RoleId *uuid.UUID `form:"role_id"`
}

func AssignUserQueryRequestDto(c *gin.Context) *UserQueryRequestDto {
	return &UserQueryRequestDto{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
