package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	userDtos "event-backend/app/user/dtos"
)

type UserServiceInterface interface {
	Create(ctx context.Context, dto *userDtos.UserCreateRequestDto) *entities.UserEntity
	Pagination(ctx context.Context, dto *userDtos.UserQueryRequestDto) *infradtos.PaginationResultDto[entities.UserEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.UserEntity
	Update(ctx context.Context, dto *userDtos.UserUpdateRequestDto) *entities.UserEntity
	Delete(ctx context.Context, id uuid.UUID)
	UpgradeToOrganizer(ctx context.Context, userId uuid.UUID)
}
