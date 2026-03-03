package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type UserQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
}

func AssignUserQueryRequestDto(c *gin.Context) *UserQueryRequestDto {
	return &UserQueryRequestDto{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
