package dtos

import (
	"github.com/google/uuid"
)

type EventRegistrationCancelRequestDto struct {
	Id     			uuid.UUID `json:"id" binding:"required"`
	CurrentUserId 	*uuid.UUID `json:"-"`
}
