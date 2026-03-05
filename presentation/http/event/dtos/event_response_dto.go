package dtos

import (
	"time"

	"event-backend/app/event/enums"
	"event-backend/entities"

	"github.com/google/uuid"
)

type EventResponseDto struct {
	Id              uuid.UUID `json:"id"`
	OrganizerUserId uuid.UUID `json:"organizer_user_id"`
	CategoryId      uuid.UUID `json:"category_id"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	Description     *string   `json:"description,omitempty"`
	Location        *string   `json:"location,omitempty"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Status          int       `json:"status"`
	StatusLabel     string    `json:"status_label"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func EventResponseDtoFromEntities(entities []*entities.EventEntity) []EventResponseDto {
	totalEvents := len(entities)
	eventResponses := make([]EventResponseDto, totalEvents)
	for i, event := range entities {
		eventResponses[i] = EventResponseDto{
			Id:              event.Id,
			OrganizerUserId: event.OrganizerUserId,
			CategoryId:      event.CategoryId,
			Title:           event.Title,
			Slug:            event.Slug,
			Description:     event.Description,
			Location:        event.Location,
			StartDate:       event.StartDate,
			EndDate:         event.EndDate,
			Status:          event.Status,
			StatusLabel:     enums.EventStatusEnum(event.Status).GetLabel(),
			CreatedAt:       event.CreatedAt,
			UpdatedAt:       event.UpdatedAt,
		}
	}

	return eventResponses
}
