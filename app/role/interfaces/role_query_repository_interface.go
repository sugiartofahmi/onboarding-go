package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	roleDtos "event-backend/app/role/dtos"
)

type RoleQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *roleDtos.RoleQueryRequestDto) *infradtos.PaginationResultDto[entities.RoleEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.RoleEntity
	FindOneByName(ctx context.Context, name string) *entities.RoleEntity
	IsExistsById(ctx context.Context, id uuid.UUID) bool
	IsExistsByName(ctx context.Context, name string) bool
	IsExistsByNameExcludeId(ctx context.Context, name string, excludeID uuid.UUID) bool
}
