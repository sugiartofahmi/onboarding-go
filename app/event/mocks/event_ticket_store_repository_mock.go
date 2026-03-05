package mocks

import (
	"context"

	"github.com/google/uuid"

	appInterfaces "event-backend/app/event/interfaces"
	"event-backend/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type EventTicketStoreRepositoryMock struct {
	mock.Mock
}

func (m *EventTicketStoreRepositoryMock) Create(ctx context.Context, entity *entities.EventTicketEntity) *entities.EventTicketEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventTicketEntity)
}

func (m *EventTicketStoreRepositoryMock) Update(ctx context.Context, entity *entities.EventTicketEntity) *entities.EventTicketEntity {
	args := m.Called(ctx, entity)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventTicketEntity)
}

func (m *EventTicketStoreRepositoryMock) Delete(ctx context.Context, id uuid.UUID) bool {
	args := m.Called(ctx, id)
	return args.Bool(0)
}

func (m *EventTicketStoreRepositoryMock) BulkCreate(ctx context.Context, entitiesArg *[]entities.EventTicketEntity) *[]entities.EventTicketEntity {
	args := m.Called(ctx, entitiesArg)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*[]entities.EventTicketEntity)
}

func (m *EventTicketStoreRepositoryMock) BulkDelete(ctx context.Context, entitiesArg *[]entities.EventTicketEntity) {
	m.Called(ctx, entitiesArg)
}

func (m *EventTicketStoreRepositoryMock) WithTransaction(tx *gorm.DB) appInterfaces.EventTicketStoreRepositoryInterface {
	args := m.Called(tx)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(appInterfaces.EventTicketStoreRepositoryInterface)
}
