package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
)

type EventTicketQueryRepositoryInterface interface {
	FindManyByEventId(ctx context.Context, eventId uuid.UUID) []entities.EventTicketEntity
	FindOneById(ctx context.Context, id uuid.UUID) *entities.EventTicketEntity
	FindOneByEventIdAndType(ctx context.Context, eventId uuid.UUID, ticketType int) *entities.EventTicketEntity
	FindByIdForCreateRegistration(ctx context.Context, id uuid.UUID) *entities.EventTicketEntity
}
