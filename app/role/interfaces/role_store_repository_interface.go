package interfaces

import (
	"context"

	"event-backend/entities"
)

type RoleStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity
	Update(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity
}
