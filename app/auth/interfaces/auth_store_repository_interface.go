package interfaces

import (
	"context"
	"event-backend/entities"
)

type AuthStoreRepositoryInterface interface {
	Create(ctx context.Context, user *entities.UserEntity) *entities.UserEntity
}
