package services

import (
	"context"
	"time"

	categoryInterfaces "event-backend/app/category/interfaces"
	eventConstants "event-backend/app/event/constants"
	eventInterfaces "event-backend/app/event/interfaces"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/exceptions"
	eventDtos "event-backend/presentation/http/event/dtos"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventService struct {
	db                         *gorm.DB
	eventQueryRepository       eventInterfaces.EventQueryRepositoryInterface
	eventStoreRepository       eventInterfaces.EventStoreRepositoryInterface
	eventTicketQueryRepository eventInterfaces.EventTicketQueryRepositoryInterface
	eventTicketStoreRepository eventInterfaces.EventTicketStoreRepositoryInterface
	categoryQueryRepository    categoryInterfaces.CategoryQueryRepositoryInterface
}

func NewEventService(
	db *gorm.DB,
	eventQueryRepository eventInterfaces.EventQueryRepositoryInterface,
	eventStoreRepository eventInterfaces.EventStoreRepositoryInterface,
	eventTicketQueryRepository eventInterfaces.EventTicketQueryRepositoryInterface,
	eventTicketStoreRepository eventInterfaces.EventTicketStoreRepositoryInterface,
	categoryQueryRepository categoryInterfaces.CategoryQueryRepositoryInterface,
) *EventService {
	return &EventService{
		db:                         db,
		eventQueryRepository:       eventQueryRepository,
		eventStoreRepository:       eventStoreRepository,
		eventTicketQueryRepository: eventTicketQueryRepository,
		eventTicketStoreRepository: eventTicketStoreRepository,
		categoryQueryRepository:    categoryQueryRepository,
	}
}

func (service *EventService) Pagination(ctx context.Context, dto *eventDtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity] {
	return service.eventQueryRepository.Pagination(ctx, dto)
}

func (service *EventService) Detail(ctx context.Context, id uuid.UUID) *entities.EventEntity {
	data := service.eventQueryRepository.FindOneByIdWithTickets(ctx, id)
	if data == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_NOT_FOUND))
	}

	return data
}

func (service *EventService) Create(ctx context.Context, dto *eventDtos.EventCreateRequestDto) *entities.EventEntity {
	categoryExists := service.categoryQueryRepository.FindOneById(ctx, dto.CategoryId)
	if categoryExists == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_CATEGORY_NOT_FOUND))
	}

	newEvent := dto.ToEntity()

	isExistsByTitle := service.eventQueryRepository.IsExistsByTitle(ctx, newEvent.Title)
	if isExistsByTitle {
		panic(*exceptions.UnprocessableEntityException(eventConstants.EVENT_TITLE_EXISTS))
	}

	var createdEvent *entities.EventEntity
	service.db.Transaction(func(tx *gorm.DB) error {
		createdEvent = service.eventStoreRepository.WithTransaction(tx).Create(ctx, newEvent)

		newTickets := dto.ToTicketEntities(createdEvent.Id)
		service.eventTicketStoreRepository.WithTransaction(tx).BulkCreate(ctx, &newTickets)
		return nil
	})

	return createdEvent
}

func (service *EventService) Update(ctx context.Context, dto *eventDtos.EventUpdateRequestDto) *entities.EventEntity {
	existingEvent := service.eventQueryRepository.FindOneById(ctx, dto.Id)
	if existingEvent == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_NOT_FOUND))
	}


	if existingEvent.OrganizerUserId != *dto.UpdatedBy {
		panic(*exceptions.ForbiddenException(eventConstants.EVENT_NOT_OWNER))
	}

	categoryExists := service.categoryQueryRepository.FindOneById(ctx, dto.CategoryId)
	if categoryExists == nil {
		panic(*exceptions.BadRequestException(eventConstants.EVENT_CATEGORY_NOT_FOUND))
	}

	updateEvent := dto.ToEntity(existingEvent)

	isExistsByTitle := service.eventQueryRepository.IsExistsByTitleExcludeId(ctx, updateEvent.Title, updateEvent.Id)
	if isExistsByTitle {
		panic(*exceptions.BadRequestException(eventConstants.EVENT_TITLE_EXISTS))
	}

	existingEventTicket := service.eventTicketQueryRepository.FindManyByEventId(ctx, dto.Id)

	var updatedEvent *entities.EventEntity
	service.db.Transaction(func(tx *gorm.DB) error {
		if len(existingEventTicket) > 0 {
			service.eventTicketStoreRepository.BulkDelete(ctx, &existingEventTicket)
		}

		updatedEvent = service.eventStoreRepository.WithTransaction(tx).Update(ctx, updateEvent)

		newTickets := dto.ToTicketEntities(updatedEvent.Id)
		service.eventTicketStoreRepository.WithTransaction(tx).BulkCreate(ctx, &newTickets)

		return nil
	})

	return updatedEvent
}

func (service *EventService) SoftDelete(ctx context.Context, id uuid.UUID, currentUserId uuid.UUID) {
	event := service.eventQueryRepository.FindOneById(ctx, id)
	if event == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_NOT_FOUND))
	}

	if event.OrganizerUserId != currentUserId {
		panic(*exceptions.ForbiddenException(eventConstants.EVENT_NOT_OWNER))
	}

	now := time.Now()
	event.DeletedAt = gorm.DeletedAt{
		Time:  now,
		Valid: true,
	}

	service.eventStoreRepository.Update(ctx, event)
}
