package dtos

import (
	"time"

	eventregistrationenums "event-backend/app/event_registration/enums"
	"event-backend/entities"

	"github.com/google/uuid"
)

type EventRegistrationResponseDto struct {
	Id            uuid.UUID `json:"id"`
	UserId        uuid.UUID `json:"user_id"`
	EventTicketId uuid.UUID `json:"event_ticket_id"`
	EventId       uuid.UUID `json:"event_id"`
	Status        int       `json:"status"`
	StatusLabel   string    `json:"status_label"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func EventRegistrationResponseDtoFromEntities(entities []*entities.EventRegistrationEntity) []EventRegistrationResponseDto {
	responses := make([]EventRegistrationResponseDto, len(entities))
	for i, entity := range entities {
		responses[i] = EventRegistrationResponseDto{
			Id:            entity.Id,
			UserId:        entity.UserId,
			EventTicketId: entity.EventTicketId,
			EventId:       entity.EventId,
			Status:        entity.Status,
			StatusLabel:   eventregistrationenums.EventRegistrationStatusEnum(entity.Status).GetLabel(),
			CreatedAt:     entity.CreatedAt,
			UpdatedAt:     entity.UpdatedAt,
		}
	}
	return responses
}
