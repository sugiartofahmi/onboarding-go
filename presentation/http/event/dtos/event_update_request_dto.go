package dtos

import (
	"time"

	"github.com/google/uuid"

	"event-backend/entities"
	"event-backend/infrastructure/utils"
)

type EventTicketUpdateRequestDto struct {
	Type  int       `json:"type" binding:"required,min=1"`
	Price float64   `json:"price" binding:"required,min=0"`
	Quota int       `json:"quota" binding:"required,min=1"`
}


type EventUpdateRequestDto struct {
	Title       string                        `json:"title" binding:"required,min=3,max=255"`
	CategoryId  uuid.UUID                     `json:"category_id" binding:"required"`
	Description *string                       `json:"description"`
	Location    *string                       `json:"location"`
	StartDate   time.Time                     `json:"start_date" binding:"required"`
	EndDate     time.Time                     `json:"end_date" binding:"required"`
	Tickets     []EventTicketUpdateRequestDto `json:"tickets" binding:"required,min=1,dive"`
	Id          uuid.UUID                     `json:"-"`
}

func (dto *EventUpdateRequestDto) ToEntity(existingEntity *entities.EventEntity) *entities.EventEntity {
	existingEntity.Title = dto.Title
	existingEntity.Slug = utils.GenerateSlug(dto.Title)
	existingEntity.CategoryId = dto.CategoryId
	existingEntity.Description = dto.Description
	existingEntity.Location = dto.Location
	existingEntity.StartDate = dto.StartDate
	existingEntity.EndDate = dto.EndDate
	return existingEntity
}

func (dto *EventUpdateRequestDto) ToTicketEntities(eventId uuid.UUID) []entities.EventTicketEntity {
	tickets := make([]entities.EventTicketEntity, len(dto.Tickets))
	for i, dto := range dto.Tickets {
		tickets[i] = entities.EventTicketEntity{
			EventId: eventId,
			Type:    dto.Type,
			Price:   dto.Price,
			Quota:   dto.Quota,
		}
	}
	return tickets
}

