package mocks

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	categorydtos "event-backend/presentation/http/category/dtos"
	"github.com/stretchr/testify/mock"
)

type CategoryQueryRepositoryMock struct {
	mock.Mock
}

func (m *CategoryQueryRepositoryMock) Pagination(ctx context.Context, dto *categorydtos.CategoryQueryRequestDto) *infradtos.PaginationResultDto[entities.CategoryEntity] {
	args := m.Called(ctx, dto)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*infradtos.PaginationResultDto[entities.CategoryEntity])
}

func (m *CategoryQueryRepositoryMock) FindOneById(ctx context.Context, id uuid.UUID) *entities.CategoryEntity {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.CategoryEntity)
}

func (m *CategoryQueryRepositoryMock) FindOneBySlug(ctx context.Context, slug string) *entities.CategoryEntity {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.CategoryEntity)
}

func (m *CategoryQueryRepositoryMock) IsExistsByName(ctx context.Context, name string) bool {
	args := m.Called(ctx, name)
	return args.Bool(0)
}

func (m *CategoryQueryRepositoryMock) IsExistsByNameExcludeId(ctx context.Context, name string, excludeID uuid.UUID) bool {
	args := m.Called(ctx, name, excludeID)
	return args.Bool(0)
}

func (m *CategoryQueryRepositoryMock) IsExistsBySlug(ctx context.Context, slug string) bool {
	args := m.Called(ctx, slug)
	return args.Bool(0)
}

func (m *CategoryQueryRepositoryMock) IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool {
	args := m.Called(ctx, slug, excludeID)
	return args.Bool(0)
}
