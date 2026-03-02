package interfaces

import (
	"context"

	"github.com/google/uuid"

	categorydtos "event-backend/app/category/dtos"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
)

type CategoryQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *categorydtos.CategoryQueryRequestDTO) *infradtos.PaginationResultDto[entities.CategoryEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.CategoryEntity
	FindOneBySlug(ctx context.Context, slug string) *entities.CategoryEntity
	IsExistsByName(ctx context.Context, name string) bool
	IsExistsByNameExcludeId(ctx context.Context, name string, excludeID uuid.UUID) bool
	IsExistsBySlug(ctx context.Context, slug string) bool
	IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool
}
