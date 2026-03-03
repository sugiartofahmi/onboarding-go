package interfaces

import (
	"context"

	"event-backend/entities"
)

type RoleQueryRepositoryInterface interface {
	FindOneByName(ctx context.Context, name string) *entities.RoleEntity
}
