package interfaces

import (
	"context"

	"event-backend/entities"
)

type CategoryStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity
	Update(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity
}
