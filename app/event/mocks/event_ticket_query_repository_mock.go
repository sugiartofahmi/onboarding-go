package mocks

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	"github.com/stretchr/testify/mock"
)

type EventTicketQueryRepositoryMock struct {
	mock.Mock
}

func (m *EventTicketQueryRepositoryMock) FindManyByEventId(ctx context.Context, eventId uuid.UUID) []entities.EventTicketEntity {
	args := m.Called(ctx, eventId)
	return args.Get(0).([]entities.EventTicketEntity)
}

func (m *EventTicketQueryRepositoryMock) FindOneById(ctx context.Context, id uuid.UUID) *entities.EventTicketEntity {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventTicketEntity)
}

func (m *EventTicketQueryRepositoryMock) FindOneByEventIdAndType(ctx context.Context, eventId uuid.UUID, ticketType int) *entities.EventTicketEntity {
	args := m.Called(ctx, eventId, ticketType)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventTicketEntity)
}
