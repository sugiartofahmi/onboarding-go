package dtos

import (
	"github.com/google/uuid"

	"event-backend/entities"
)

type EventRegistrationCreateRequestDto struct {
	EventTicketId 	uuid.UUID `json:"event_ticket_id" binding:"required"`
	CurrentUserId	*uuid.UUID `json:"-"`
}

func (dto *EventRegistrationCreateRequestDto) ToEntity() *entities.EventRegistrationEntity {
	return &entities.EventRegistrationEntity{
		EventTicketId: dto.EventTicketId,
	}
}

