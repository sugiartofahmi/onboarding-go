package mocks

import (
	"context"

	"event-backend/entities"
	"github.com/stretchr/testify/mock"
)

type UserStoreRepositoryMock struct {
	mock.Mock
}

func (m *UserStoreRepositoryMock) Create(ctx context.Context, entity *entities.UserEntity) *entities.UserEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}

func (m *UserStoreRepositoryMock) Update(ctx context.Context, entity *entities.UserEntity) *entities.UserEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}
