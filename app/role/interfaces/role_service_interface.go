package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	roledtos "event-backend/presentation/http/role/dtos"
)

type RoleServiceInterface interface {
	Pagination(ctx context.Context, dto *roledtos.RoleQueryRequestDto) *infradtos.PaginationResultDto[entities.RoleEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.RoleEntity
	Create(ctx context.Context, dto *roledtos.RoleCreateRequestDto) *entities.RoleEntity
	Update(ctx context.Context, dto *roledtos.RoleUpdateRequestDto) *entities.RoleEntity
	SoftDelete(ctx context.Context, id uuid.UUID)
}
