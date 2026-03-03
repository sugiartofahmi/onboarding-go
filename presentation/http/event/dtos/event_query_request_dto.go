package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type EventQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
}

func AssignEventQueryRequestDto(c *gin.Context) *EventQueryRequestDto {
	return &EventQueryRequestDto{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}
