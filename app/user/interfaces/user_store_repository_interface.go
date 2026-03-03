package interfaces

import (
	"context"
	"event-backend/entities"
)

type UserStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.UserEntity) *entities.UserEntity
	Update(ctx context.Context, entity *entities.UserEntity) *entities.UserEntity
}