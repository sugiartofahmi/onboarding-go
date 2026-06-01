package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
)

type CategoryStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity
	Update(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity
	DeleteById(ctx context.Context, id uuid.UUID) error
}
