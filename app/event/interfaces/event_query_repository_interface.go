package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventdtos "event-backend/presentation/http/event/dtos"
)

type EventQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *eventdtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.EventEntity
	FindOneBySlug(ctx context.Context, slug string) *entities.EventEntity
	IsExistsByTitle(ctx context.Context, title string) bool
	IsExistsByTitleExcludeId(ctx context.Context, title string, excludeID uuid.UUID) bool
	IsExistsBySlug(ctx context.Context, slug string) bool
	IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool
	IsExistsByCategoryId(ctx context.Context, categoryId uuid.UUID) bool
}
