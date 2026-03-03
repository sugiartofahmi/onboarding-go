package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type UserQueryRequestDTO struct {
	infradtos.PaginationQueryRequestDto
}

func AssignUserQueryRequestDTO(c *gin.Context) *UserQueryRequestDTO {
	return &UserQueryRequestDTO{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
