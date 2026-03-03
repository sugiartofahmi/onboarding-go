package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	categorydtos "event-backend/presentation/http/category/dtos"
)

type CategoryServiceInterface interface {
	Pagination(ctx context.Context, dto *categorydtos.CategoryQueryRequestDto) *infradtos.PaginationResultDto[entities.CategoryEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.CategoryEntity
	Create(ctx context.Context, entity *categorydtos.CategoryCreateRequestDto) *entities.CategoryEntity
	Update(ctx context.Context, entity *categorydtos.CategoryUpdateRequestDto) *entities.CategoryEntity
	SoftDelete(ctx context.Context, id uuid.UUID)
}
