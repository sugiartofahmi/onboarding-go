package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventregistrationdtos "event-backend/presentation/http/event_registration/dtos"
)

type EventRegistrationServiceInterface interface {
	Pagination(ctx context.Context, dto *eventregistrationdtos.EventRegistrationQueryRequestDto) *infradtos.PaginationResultDto[entities.EventRegistrationEntity]
	Create(ctx context.Context, dto *eventregistrationdtos.EventRegistrationCreateRequestDto) *entities.EventRegistrationEntity
	Cancel(ctx context.Context, dto *eventregistrationdtos.EventRegistrationCancelRequestDto) *entities.EventRegistrationEntity
	Detail(ctx context.Context, id uuid.UUID) *entities.EventRegistrationEntity
}
