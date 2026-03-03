package interfaces

import (
	"context"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	userdtos "event-backend/presentation/http/user/dtos"

	"github.com/google/uuid"
)

type UserQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *userdtos.UserQueryRequestDto) *infradtos.PaginationResultDto[entities.UserEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.UserEntity
	FindOneByIdWithRole(ctx context.Context, id uuid.UUID) *entities.UserEntity
	FindOneByEmail(ctx context.Context, email string) *entities.UserEntity
	IsExistsByEmail(ctx context.Context, email string) bool
	IsExistsByEmailExcludeId(ctx context.Context, email string, excludeID uuid.UUID) bool
}