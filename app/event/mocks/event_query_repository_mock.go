package mocks

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventdtos "event-backend/presentation/http/event/dtos"
	"github.com/stretchr/testify/mock"
)

type EventQueryRepositoryMock struct {
	mock.Mock
}

func (m *EventQueryRepositoryMock) Pagination(ctx context.Context, dto *eventdtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity] {
	args := m.Called(ctx, dto)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*infradtos.PaginationResultDto[entities.EventEntity])
}

func (m *EventQueryRepositoryMock) FindOneById(ctx context.Context, id uuid.UUID) *entities.EventEntity {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventEntity)
}

func (m *EventQueryRepositoryMock) FindOneBySlug(ctx context.Context, slug string) *entities.EventEntity {
	args := m.Called(ctx, slug)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventEntity)
}

func (m *EventQueryRepositoryMock) FindOneByIdWithTickets(ctx context.Context, id uuid.UUID) *entities.EventEntity {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*entities.EventEntity)
}

func (m *EventQueryRepositoryMock) IsExistsByTitle(ctx context.Context, title string) bool {
	args := m.Called(ctx, title)
	return args.Bool(0)
}

func (m *EventQueryRepositoryMock) IsExistsByTitleExcludeId(ctx context.Context, title string, excludeID uuid.UUID) bool {
	args := m.Called(ctx, title, excludeID)
	return args.Bool(0)
}

func (m *EventQueryRepositoryMock) IsExistsBySlug(ctx context.Context, slug string) bool {
	args := m.Called(ctx, slug)
	return args.Bool(0)
}

func (m *EventQueryRepositoryMock) IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool {
	args := m.Called(ctx, slug, excludeID)
	return args.Bool(0)
}

func (m *EventQueryRepositoryMock) IsExistsByCategoryId(ctx context.Context, categoryId uuid.UUID) bool {
	args := m.Called(ctx, categoryId)
	return args.Bool(0)
}
