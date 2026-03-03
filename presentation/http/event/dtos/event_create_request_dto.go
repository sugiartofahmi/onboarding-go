package dtos

import (
	"time"

	"event-backend/entities"
	"event-backend/infrastructure/utils"

	"github.com/google/uuid"
)

type EventCreateRequestDto struct {
	Title       string    `json:"title" binding:"not_empty,min=3,max=255"`
	CategoryId  uuid.UUID `json:"category_id" binding:"required"`
	Description *string   `json:"description"`
	Location    *string   `json:"location"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
}

func (dto *EventCreateRequestDto) ToEntity(organizerUserId uuid.UUID) *entities.EventEntity {
	return &entities.EventEntity{
		OrganizerUserId: organizerUserId,
		CategoryId:      dto.CategoryId,
		Title:           dto.Title,
		Slug:            utils.GenerateSlug(dto.Title),
		Description:     dto.Description,
		Location:        dto.Location,
		StartDate:       dto.StartDate,
		EndDate:         dto.EndDate,
		Status:          0,
	}
}
