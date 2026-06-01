package dtos

import (
	"time"

	"event-backend/entities"
	"event-backend/infrastructure/utils"

	eventStatusEnum "event-backend/app/event/enums"

	"github.com/google/uuid"
)

type EventTicketCreateRequestDto struct {
	EventId uuid.UUID `json:"event_id" binding:"required"`
	Type    int       `json:"type" binding:"required,min=1"`
	Price   float64   `json:"price" binding:"required,min=0"`
	Quota   int       `json:"quota" binding:"required,min=1"`
}

type EventCreateRequestDto struct {
	Title       string                        `json:"title" binding:"required,min=3,max=255"`
	CategoryId  uuid.UUID                     `json:"category_id" binding:"required"`
	Description *string                       `json:"description"`
	Location    *string                       `json:"location"`
	StartDate   time.Time                     `json:"start_date" binding:"required"`
	EndDate     time.Time                     `json:"end_date" binding:"required"`
	Tickets     []EventTicketCreateRequestDto `json:"tickets" binding:"required,min=1,dive"`
	CreatedBy   *uuid.UUID					  `json:"-"`
}

func (dto *EventCreateRequestDto) ToEntity() *entities.EventEntity {

	return &entities.EventEntity{
		CategoryId:      dto.CategoryId,
		Title:           dto.Title,
		Slug:            utils.GenerateSlug(dto.Title),
		Description:     dto.Description,
		Location:        dto.Location,
		StartDate:       dto.StartDate,
		EndDate:         dto.EndDate,
		Status:          int(eventStatusEnum.Draft), // Default status is Draft when creating a new event
		CreatedBy:       dto.CreatedBy,
	}
}

func (dto *EventCreateRequestDto) ToTicketEntities(eventId uuid.UUID) []entities.EventTicketEntity {
	dtoTickets := dto.Tickets
	tickets := make([]entities.EventTicketEntity, len(dtoTickets))
	for i, dtoTicket := range dtoTickets {
		tickets[i] = entities.EventTicketEntity{
			EventId: eventId,
			Type:    dtoTicket.Type,
			Price:   dtoTicket.Price,
			Quota:   dtoTicket.Quota,
			CreatedBy: dto.CreatedBy,
		}
	}
	return tickets
}
