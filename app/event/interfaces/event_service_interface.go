package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventDtos "event-backend/app/event/dtos"
)

type EventServiceInterface interface {
	Pagination(ctx context.Context, dto *eventDtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.EventEntity
	Create(ctx context.Context, dto *eventDtos.EventCreateRequestDto) *entities.EventEntity
	Update(ctx context.Context, dto *eventDtos.EventUpdateRequestDto) *entities.EventEntity
	SoftDelete(ctx context.Context, id uuid.UUID, currentUserId uuid.UUID)
}
