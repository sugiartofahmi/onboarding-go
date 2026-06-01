package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type CategoryQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
}

func AssignCategoryQueryRequestDto(c *gin.Context) *CategoryQueryRequestDto {
	return &CategoryQueryRequestDto{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
