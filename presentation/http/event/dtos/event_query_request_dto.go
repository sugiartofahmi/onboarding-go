package dtos

import (
	"github.com/gin-gonic/gin"

	infradtos "event-backend/infrastructure/dtos"
)

type EventQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
}

func AssignEventQueryRequestDto(c *gin.Context) *EventQueryRequestDto {
	q := &EventQueryRequestDto{}
	c.ShouldBindQuery(q)
	return q
}
