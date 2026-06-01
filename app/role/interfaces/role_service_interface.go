package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	roleDtos "event-backend/app/role/dtos"
)

type RoleServiceInterface interface {
	Pagination(ctx context.Context, dto *roleDtos.RoleQueryRequestDto) *infradtos.PaginationResultDto[entities.RoleEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.RoleEntity
	Create(ctx context.Context, dto *roleDtos.RoleCreateRequestDto) *entities.RoleEntity
	Update(ctx context.Context, dto *roleDtos.RoleUpdateRequestDto) *entities.RoleEntity
	SoftDelete(ctx context.Context, id uuid.UUID)
	Delete(ctx context.Context, id uuid.UUID)
}
