package interfaces

import (
	"context"

	"event-backend/entities"

	"github.com/google/uuid"
)

type RoleQueryRepositoryInterface interface {
	FindOneById(ctx context.Context, id uuid.UUID) *entities.RoleEntity
	FindOneByName(ctx context.Context, name string) *entities.RoleEntity
	IsExistsById(ctx context.Context, id uuid.UUID) bool
}
