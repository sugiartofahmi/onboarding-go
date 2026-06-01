package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	categoryDtos "event-backend/app/category/dtos"
)

type CategoryServiceInterface interface {
	Pagination(ctx context.Context, dto *categoryDtos.CategoryQueryRequestDto) *infradtos.PaginationResultDto[entities.CategoryEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.CategoryEntity
	Create(ctx context.Context, entity *categoryDtos.CategoryCreateRequestDto) *entities.CategoryEntity
	Update(ctx context.Context, entity *categoryDtos.CategoryUpdateRequestDto) *entities.CategoryEntity
	SoftDelete(ctx context.Context, id uuid.UUID)
}
