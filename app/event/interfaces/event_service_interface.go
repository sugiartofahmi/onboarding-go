package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventdtos "event-backend/presentation/http/event/dtos"
)

type EventServiceInterface interface {
	Pagination(ctx context.Context, dto *eventdtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity]
	Detail(ctx context.Context, id uuid.UUID) *entities.EventEntity
	Create(ctx context.Context, dto *eventdtos.EventCreateRequestDto) *entities.EventEntity
	Update(ctx context.Context, dto *eventdtos.EventUpdateRequestDto) *entities.EventEntity
	SoftDelete(ctx context.Context, id uuid.UUID, currentUserId uuid.UUID)
}
