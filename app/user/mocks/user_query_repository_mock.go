package mocks

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	userdtos "event-backend/presentation/http/user/dtos"
	"github.com/stretchr/testify/mock"
)

type UserQueryRepositoryMock struct {
	mock.Mock
}

func (m *UserQueryRepositoryMock) Pagination(ctx context.Context, dto *userdtos.UserQueryRequestDto) *infradtos.PaginationResultDto[entities.UserEntity] {
	args := m.Called(ctx, dto)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*infradtos.PaginationResultDto[entities.UserEntity])
}

func (m *UserQueryRepositoryMock) FindOneById(ctx context.Context, id uuid.UUID) *entities.UserEntity {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}

func (m *UserQueryRepositoryMock) FindOneByIdWithRole(ctx context.Context, id uuid.UUID) *entities.UserEntity {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}

func (m *UserQueryRepositoryMock) FindOneByEmail(ctx context.Context, email string) *entities.UserEntity {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}

func (m *UserQueryRepositoryMock) IsExistsByEmail(ctx context.Context, email string) bool {
	args := m.Called(ctx, email)
	return args.Bool(0)
}

func (m *UserQueryRepositoryMock) IsExistsByEmailExcludeId(ctx context.Context, email string, excludeID uuid.UUID) bool {
	args := m.Called(ctx, email, excludeID)
	return args.Bool(0)
}
