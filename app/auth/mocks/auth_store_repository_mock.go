package mocks

import (
	"context"
	"event-backend/entities"
	"github.com/stretchr/testify/mock"
)

type AuthStoreRepositoryMock struct {
	mock.Mock
}

func (m *AuthStoreRepositoryMock) Create(ctx context.Context, user *entities.UserEntity) *entities.UserEntity {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}
