package interfaces

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"event-backend/entities"
)

type EventStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity
	Update(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity
	DeleteById(ctx context.Context, id uuid.UUID) error
	WithTransaction(tx *gorm.DB) EventStoreRepositoryInterface
}
