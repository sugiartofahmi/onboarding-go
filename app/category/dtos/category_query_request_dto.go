package dtos

import (
	"github.com/gin-gonic/gin"
	infradtos "event-backend/infrastructure/dtos"
)

type CategoryQueryRequestDTO struct {
	infradtos.PaginationQueryRequestDto
}

// AssignCategoryQueryRequestDTO assigns query params from gin context
func AssignCategoryQueryRequestDTO(c *gin.Context) *CategoryQueryRequestDTO {
	return &CategoryQueryRequestDTO{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
