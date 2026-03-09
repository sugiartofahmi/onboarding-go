package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	userdtos "event-backend/presentation/http/user/dtos"
)

type UserServiceInterface interface {
	Create(ctx context.Context, dto *userdtos.UserCreateRequestDto) *entities.UserEntity
	Pagination(ctx context.Context, dto *userdtos.UserQueryRequestDto) *infradtos.PaginationResultDto[entities.UserEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.UserEntity
	Update(ctx context.Context, dto *userdtos.UserUpdateRequestDto) *entities.UserEntity
	Delete(ctx context.Context, id uuid.UUID)
	UpgradeToOrganizer(ctx context.Context, userId uuid.UUID)
}
