package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventregistrationdtos "event-backend/presentation/http/event_registration/dtos"
)

type EventRegistrationQueryRepositoryInterface interface {
	Pagination(ctx context.Context, dto *eventregistrationdtos.EventRegistrationQueryRequestDto) *infradtos.PaginationResultDto[entities.EventRegistrationEntity]
	FindOneById(ctx context.Context, id uuid.UUID) *entities.EventRegistrationEntity
}
