package interfaces

import (
	"context"
	"event-backend/entities"
)

type AuthQueryRepositoryInterface interface {
	FindOneByEmail(ctx context.Context, email string) *entities.UserEntity
	FindOneByEmailWithRole(ctx context.Context, email string) *entities.UserEntity
	IsExistsByEmail(ctx context.Context, email string) bool
}