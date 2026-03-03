package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	roledtos "event-backend/presentation/http/role/dtos"
)

type RoleServiceInterface interface {
	Pagination(ctx context.Context, dto *roledtos.RoleQueryRequestDTO) *infradtos.PaginationResultDto[entities.RoleEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.RoleEntity
	Create(ctx context.Context, dto *roledtos.RoleCreateRequestDTO) *entities.RoleEntity
	Update(ctx context.Context, dto *roledtos.RoleUpdateRequestDTO) *entities.RoleEntity
	SoftDelete(ctx context.Context, id uuid.UUID)
}
