package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
)

type RoleStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity
	Update(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity
	DeleteById(ctx context.Context, id uuid.UUID) error
}
