package interfaces

import (
	"context"

	"event-backend/entities"
)

type EventStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity
	Update(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity
}
