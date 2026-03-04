package interfaces

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"event-backend/entities"
)

type EventTicketStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.EventTicketEntity) *entities.EventTicketEntity
	Update(ctx context.Context, entity *entities.EventTicketEntity) *entities.EventTicketEntity
	Delete(ctx context.Context, id uuid.UUID) bool
	BulkCreate(ctx context.Context, entities *[]entities.EventTicketEntity) *[]entities.EventTicketEntity
	BulkDelete(ctx context.Context, entities *[]entities.EventTicketEntity)
	WithTransaction(tx *gorm.DB) EventTicketStoreRepositoryInterface
}