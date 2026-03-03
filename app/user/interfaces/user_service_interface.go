package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	userdtos "event-backend/presentation/http/user/dtos"
)

type UserServiceInterface interface {
	Create(ctx context.Context, dto *userdtos.UserCreateRequestDTO) *entities.UserEntity
	Pagination(ctx context.Context, dto *userdtos.UserQueryRequestDTO) *infradtos.PaginationResultDto[entities.UserEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.UserEntity
	Update(ctx context.Context, dto *userdtos.UserUpdateRequestDTO) *entities.UserEntity
	Delete(ctx context.Context, id uuid.UUID)
}
