package mocks

import (
	"context"

	appInterfaces "event-backend/app/event/interfaces"
	"event-backend/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type EventStoreRepositoryMock struct {
	mock.Mock
}

func (m *EventStoreRepositoryMock) Create(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventEntity)
}

func (m *EventStoreRepositoryMock) Update(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventEntity)
}

func (m *EventStoreRepositoryMock) WithTransaction(tx *gorm.DB) appInterfaces.EventStoreRepositoryInterface {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(appInterfaces.EventStoreRepositoryInterface)
}
