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
	findCategoryCh := make(chan *entities.CategoryEntity, 1)
	isTitleExistsCh := make(chan bool, 1)

	go func() {
		findCategoryCh <- service.categoryQueryRepository.FindOneById(ctx, dto.CategoryId)
	}()

	go func() {
		isTitleExistsCh <- service.eventQueryRepository.IsExistsByTitle(ctx, dto.Title)
	}()

	findCategoryResult := <-findCategoryCh
	isTitleExistsResult := <-isTitleExistsCh

	if findCategoryResult == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_CATEGORY_NOT_FOUND))
	}
	if isTitleExistsResult {
		panic(*exceptions.UnprocessableEntityException(eventConstants.EVENT_TITLE_EXISTS))
	}

	newEvent := dto.ToEntity()

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
	findEventCh := make(chan *entities.EventEntity, 1)
	findCategoryCh := make(chan *entities.CategoryEntity, 1)
	isTitleExistsCh := make(chan bool, 1)
	findTicketsCh := make(chan []entities.EventTicketEntity, 1)

	go func() {
		findEventCh <- service.eventQueryRepository.FindOneById(ctx, dto.Id)
	}()

	go func() {
		findCategoryCh <- service.categoryQueryRepository.FindOneById(ctx, dto.CategoryId)
	}()

	go func() {
		isTitleExistsCh <- service.eventQueryRepository.IsExistsByTitleExcludeId(ctx, dto.Title, dto.Id)
	}()

	go func() {
		findTicketsCh <- service.eventTicketQueryRepository.FindManyByEventId(ctx, dto.Id)
	}()

	findEventResult := <-findEventCh
	findingCategoryResult := <-findCategoryCh
	isTitleExistsResult := <-isTitleExistsCh
	findTicketsResult := <-findTicketsCh

	if findEventResult == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_NOT_FOUND))
	}

	if findEventResult.OrganizerUserId != *dto.UpdatedBy {
		panic(*exceptions.ForbiddenException(eventConstants.EVENT_NOT_OWNER))
	}

	if findingCategoryResult == nil {
		panic(*exceptions.BadRequestException(eventConstants.EVENT_CATEGORY_NOT_FOUND))
	}

	if isTitleExistsResult {
		panic(*exceptions.BadRequestException(eventConstants.EVENT_TITLE_EXISTS))
	}

	updateEvent := dto.ToEntity(findEventResult)

	var updatedEvent *entities.EventEntity
	service.db.Transaction(func(tx *gorm.DB) error {
		if len(findTicketsResult) > 0 {
			service.eventTicketStoreRepository.BulkDelete(ctx, &findTicketsResult)
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
