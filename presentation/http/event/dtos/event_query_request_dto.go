package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	infradtos "event-backend/infrastructure/dtos"
)

type EventQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
	Status		*int `form:"status"`
	CategoryId	*uuid.UUID `form:"category_id" binding:"uuid,omitempty"`
	StartDate	*string `form:"start_date" binding:"datetime=2006-01-02,omitempty"`
	EndDate		*string `form:"end_date" binding:"datetime=2006-01-02,omitempty"`
}

func AssignEventQueryRequestDto(c *gin.Context) *EventQueryRequestDto {
	q := &EventQueryRequestDto{}
	c.ShouldBindQuery(q)
	return q
}
