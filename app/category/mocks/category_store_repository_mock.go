package mocks

import (
	"context"

	"event-backend/entities"
	"github.com/stretchr/testify/mock"
)

type CategoryStoreRepositoryMock struct {
	mock.Mock
}

func (m *CategoryStoreRepositoryMock) Create(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.CategoryEntity)
}

func (m *CategoryStoreRepositoryMock) Update(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.CategoryEntity)
}
