package services

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	eventInterfaces "event-backend/app/event/interfaces"
	eventregistrationConstants "event-backend/app/event_registration/constants"
	eventregistrationEnums "event-backend/app/event_registration/enums"
	eventregistrationRepoInterfaces "event-backend/app/event_registration/interfaces/repositories"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/exceptions"
	eventregistrationDtos "event-backend/app/event_registration/dtos"
)

type EventRegistrationService struct {
	db                               *gorm.DB
	eventRegistrationQueryRepository eventregistrationRepoInterfaces.EventRegistrationQueryRepositoryInterface
	eventRegistrationStoreRepository eventregistrationRepoInterfaces.EventRegistrationStoreRepositoryInterface
	eventTicketQueryRepository       eventInterfaces.EventTicketQueryRepositoryInterface
	eventTicketStoreRepository       eventInterfaces.EventTicketStoreRepositoryInterface
}

func NewEventRegistrationService(
	db *gorm.DB,
	eventRegistrationQueryRepository eventregistrationRepoInterfaces.EventRegistrationQueryRepositoryInterface,
	eventRegistrationStoreRepository eventregistrationRepoInterfaces.EventRegistrationStoreRepositoryInterface,
	eventTicketQueryRepository eventInterfaces.EventTicketQueryRepositoryInterface,
	eventTicketStoreRepository eventInterfaces.EventTicketStoreRepositoryInterface,
) *EventRegistrationService {
	return &EventRegistrationService{
		db:                               db,
		eventRegistrationQueryRepository: eventRegistrationQueryRepository,
		eventRegistrationStoreRepository: eventRegistrationStoreRepository,
		eventTicketQueryRepository:       eventTicketQueryRepository,
		eventTicketStoreRepository:       eventTicketStoreRepository,
	}
}

func (service *EventRegistrationService) Pagination(ctx context.Context, dto *eventregistrationDtos.EventRegistrationQueryRequestDto) *infradtos.PaginationResultDto[entities.EventRegistrationEntity] {
	return service.eventRegistrationQueryRepository.Pagination(ctx, dto)
}

func (service *EventRegistrationService) Detail(ctx context.Context, id uuid.UUID) *entities.EventRegistrationEntity {
	registration := service.eventRegistrationQueryRepository.FindOneById(ctx, id)
	if registration == nil {
		panic(*exceptions.NotFoundException(eventregistrationConstants.REGISTRATION_NOT_FOUND))
	}

	return registration
}

func (service *EventRegistrationService) Create(ctx context.Context, dto *eventregistrationDtos.EventRegistrationCreateRequestDto) *entities.EventRegistrationEntity {
	ticket := service.eventTicketQueryRepository.FindByIdForCreateRegistration(ctx, dto.EventTicketId)
	if ticket == nil {
		panic(*exceptions.NotFoundException(eventregistrationConstants.TICKET_NOT_FOUND))
	}

	isQuotaAvailable := ticket.RegisteredCount < ticket.Quota
	if !isQuotaAvailable {
		panic(*exceptions.BadRequestException(eventregistrationConstants.REGISTRATION_EVENT_FULL))
	}

	const ticketQuantity = 1
	entity := dto.ToEntity()
	entity.EventId = ticket.EventId
	entity.Status = int(eventregistrationEnums.Confirmed)
	entity.CreatedBy = dto.CurrentUserId

	var created *entities.EventRegistrationEntity
	service.db.Transaction(func(tx *gorm.DB) error {
		created = service.eventRegistrationStoreRepository.WithTransaction(tx).Create(ctx, entity)
		service.eventTicketStoreRepository.WithTransaction(tx).IncrementRegisteredCount(ctx, dto.EventTicketId, ticketQuantity)
		return nil
	})

	return created
}

func (service *EventRegistrationService) Cancel(ctx context.Context, dto *eventregistrationDtos.EventRegistrationCancelRequestDto) *entities.EventRegistrationEntity {
	registration := service.eventRegistrationQueryRepository.FindOneById(ctx, dto.Id)
	if registration == nil {
		panic(*exceptions.NotFoundException(eventregistrationConstants.REGISTRATION_NOT_FOUND))
	}

	isCancelled := registration.Status == int(eventregistrationEnums.Cancelled)
	if isCancelled {
		panic(*exceptions.BadRequestException(eventregistrationConstants.REGISTRATION_ALREADY_CANCELLED))
	}

	const ticketQuantity = 1
	registration.Status = int(eventregistrationEnums.Cancelled)

	var updated *entities.EventRegistrationEntity
	service.db.Transaction(func(tx *gorm.DB) error {
		updated = service.eventRegistrationStoreRepository.WithTransaction(tx).Update(ctx, registration)
		service.eventTicketStoreRepository.WithTransaction(tx).DecrementRegisteredCount(ctx, registration.EventTicketId, ticketQuantity)
		return nil
	})

	return updated
}
