package dtos

import (
	"time"

	"github.com/google/uuid"

	"event-backend/entities"
	"event-backend/infrastructure/utils"
)

type EventUpdateRequestDto struct {
	Title       string    `json:"title" binding:"not_empty,min=3,max=255"`
	CategoryId  uuid.UUID `json:"category_id" binding:"required"`
	Description *string   `json:"description"`
	Location    *string   `json:"location"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	Id          uuid.UUID `json:"-"`
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
