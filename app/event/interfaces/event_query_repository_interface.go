package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventDtos "event-backend/app/event/dtos"
)

type EventQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *eventDtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.EventEntity
	FindOneBySlug(ctx context.Context, slug string) *entities.EventEntity
	FindOneByIdWithTickets(ctx context.Context, id uuid.UUID) *entities.EventEntity
	IsExistsById(ctx context.Context, id uuid.UUID) bool
	IsExistsByTitle(ctx context.Context, title string) bool
	IsExistsByTitleExcludeId(ctx context.Context, title string, excludeID uuid.UUID) bool
	IsExistsBySlug(ctx context.Context, slug string) bool
	IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool
	IsExistsByCategoryId(ctx context.Context, categoryId uuid.UUID) bool
}
