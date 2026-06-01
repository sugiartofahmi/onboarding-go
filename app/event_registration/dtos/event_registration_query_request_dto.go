package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	infradtos "event-backend/infrastructure/dtos"
)

type EventRegistrationQueryRequestDto struct {
	infradtos.PaginationQueryRequestDto
	EventId 				string `form:"event_id"`
	Status  				string `form:"status"`
	CurrentUserId  			uuid.UUID
	CurrentUserRoleName 	string
}

func AssignEventRegistrationQueryRequestDto(c *gin.Context) *EventRegistrationQueryRequestDto {
	return &EventRegistrationQueryRequestDto{
		PaginationQueryRequestDto: *infradtos.AssignPaginationQueryRequestDto(c),
	}
}

