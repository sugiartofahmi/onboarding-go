package mocks

import (
	"context"
	"event-backend/entities"
	"github.com/stretchr/testify/mock"
)

type AuthQueryRepositoryMock struct {
	mock.Mock
}

func (m *AuthQueryRepositoryMock) FindOneByEmail(ctx context.Context, email string) *entities.UserEntity {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}

func (m *AuthQueryRepositoryMock) FindOneByEmailWithRole(ctx context.Context, email string) *entities.UserEntity {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.UserEntity)
}

func (m *AuthQueryRepositoryMock) IsExistsByEmail(ctx context.Context, email string) bool {
	args := m.Called(ctx, email)
	return args.Bool(0)
}
