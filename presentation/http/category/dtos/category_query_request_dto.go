package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type CategoryQueryRequestDTO struct {
	infradtos.PaginationQueryRequestDto
}

func AssignCategoryQueryRequestDTO(c *gin.Context) *CategoryQueryRequestDTO {
	return &CategoryQueryRequestDTO{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
