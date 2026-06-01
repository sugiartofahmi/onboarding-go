package interfaces

import (
	"context"

	"github.com/google/uuid"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	eventregistrationDtos "event-backend/app/event_registration/dtos"
)

type EventRegistrationServiceInterface interface {
	Pagination(ctx context.Context, dto *eventregistrationDtos.EventRegistrationQueryRequestDto) *infradtos.PaginationResultDto[entities.EventRegistrationEntity]
	Create(ctx context.Context, dto *eventregistrationDtos.EventRegistrationCreateRequestDto) *entities.EventRegistrationEntity
	Cancel(ctx context.Context, dto *eventregistrationDtos.EventRegistrationCancelRequestDto) *entities.EventRegistrationEntity
	Detail(ctx context.Context, id uuid.UUID) *entities.EventRegistrationEntity
}
