package dtos

import (
	"time"

	"event-backend/app/event/enums"
	"event-backend/entities"

	"github.com/google/uuid"
)

type EventDetailResponseDto struct {
	Id              uuid.UUID                `json:"id"`
	OrganizerUserId uuid.UUID                `json:"organizer_user_id"`
	CategoryId      uuid.UUID                `json:"category_id"`
	Title           string                   `json:"title"`
	Slug            string                   `json:"slug"`
	Description     *string                  `json:"description,omitempty"`
	Location        *string                  `json:"location,omitempty"`
	StartDate       time.Time                `json:"start_date"`
	EndDate         time.Time                `json:"end_date"`
	Status          int                      `json:"status"`
	StatusLabel     string                   `json:"status_label"`
	Tickets         []EventTicketResponseDto `json:"tickets,omitempty"`
	CreatedAt       time.Time                `json:"created_at"`
	UpdatedAt       time.Time                `json:"updated_at"`
}

type EventTicketResponseDto struct {
	Id              uuid.UUID `json:"id"`
	EventId         uuid.UUID `json:"event_id"`
	Type            int       `json:"type"`
	TypeLabel       string    `json:"type_label"`
	Price           float64   `json:"price"`
	Quota           int       `json:"quota"`
	RegisteredCount int       `json:"registered_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func EventDetailResponseDtoFromEntity(entity entities.EventEntity) EventDetailResponseDto {
	ticketResponses := make([]EventTicketResponseDto, len(entity.Tickets))
	for i, ticket := range entity.Tickets {
		ticketResponses[i] = EventTicketResponseDto{
			Id:              ticket.Id,
			EventId:         ticket.EventId,
			Type:            ticket.Type,
			TypeLabel:       enums.EventTicketTypeEnum(ticket.Type).GetLabel(),
			Price:           ticket.Price,
			Quota:           ticket.Quota,
			RegisteredCount: ticket.RegisteredCount,
			CreatedAt:       ticket.CreatedAt,
			UpdatedAt:       ticket.UpdatedAt,
		}
	}

	return EventDetailResponseDto{
		Id:              entity.Id,
		OrganizerUserId: entity.OrganizerUserId,
		CategoryId:      entity.CategoryId,
		Title:           entity.Title,
		Slug:            entity.Slug,
		Description:     entity.Description,
		Location:        entity.Location,
		StartDate:       entity.StartDate,
		EndDate:         entity.EndDate,
		Status:          entity.Status,
		StatusLabel:     enums.EventStatusEnum(entity.Status).GetLabel(),
		Tickets:         ticketResponses,
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
	}
}
