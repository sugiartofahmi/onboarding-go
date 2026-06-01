package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	categoryDtos "event-backend/app/category/dtos"
)

type CategoryQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *categoryDtos.CategoryQueryRequestDto) *infradtos.PaginationResultDto[entities.CategoryEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.CategoryEntity
	FindOneBySlug(ctx context.Context, slug string) *entities.CategoryEntity
	IsExistsById(ctx context.Context, id uuid.UUID) bool
	IsExistsByName(ctx context.Context, name string) bool
	IsExistsByNameExcludeId(ctx context.Context, name string, excludeID uuid.UUID) bool
	IsExistsBySlug(ctx context.Context, slug string) bool
	IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool
}
