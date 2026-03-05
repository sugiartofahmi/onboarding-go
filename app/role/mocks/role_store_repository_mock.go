package mocks

import (
	"context"

	"event-backend/entities"
	"github.com/stretchr/testify/mock"
)

type RoleStoreRepositoryMock struct {
	mock.Mock
}

func (m *RoleStoreRepositoryMock) Create(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.RoleEntity)
}

func (m *RoleStoreRepositoryMock) Update(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.RoleEntity)
}
