package interfaces

import (
	"context"

	"gorm.io/gorm"

	"event-backend/entities"
)

type EventStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity
	Update(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity
	WithTransaction(tx *gorm.DB) EventStoreRepositoryInterface
}
