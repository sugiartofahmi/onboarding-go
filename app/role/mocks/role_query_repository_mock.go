package mocks

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	roledtos "event-backend/presentation/http/role/dtos"
	"github.com/stretchr/testify/mock"
)

type RoleQueryRepositoryMock struct {
	mock.Mock
}

func (m *RoleQueryRepositoryMock) Pagination(ctx context.Context, dto *roledtos.RoleQueryRequestDto) *infradtos.PaginationResultDto[entities.RoleEntity] {
	args := m.Called(ctx, dto)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*infradtos.PaginationResultDto[entities.RoleEntity])
}

func (m *RoleQueryRepositoryMock) FindOneById(ctx context.Context, id uuid.UUID) *entities.RoleEntity {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.RoleEntity)
}

func (m *RoleQueryRepositoryMock) FindOneByName(ctx context.Context, name string) *entities.RoleEntity {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.RoleEntity)
}

func (m *RoleQueryRepositoryMock) IsExistsById(ctx context.Context, id uuid.UUID) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *RoleQueryRepositoryMock) IsExistsByName(ctx context.Context, name string) bool {
	args := m.Called(ctx, name)
	return args.Bool(0)
}

func (m *RoleQueryRepositoryMock) IsExistsByNameExcludeId(ctx context.Context, name string, excludeID uuid.UUID) bool {
	args := m.Called(ctx, name, excludeID)
	return args.Bool(0)
}
