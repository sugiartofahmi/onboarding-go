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
	eventQueryRepository    eventInterfaces.EventQueryRepositoryInterface
	eventStoreRepository    eventInterfaces.EventStoreRepositoryInterface
	categoryQueryRepository categoryInterfaces.CategoryQueryRepositoryInterface
}

func NewEventService(
	eventQueryRepository eventInterfaces.EventQueryRepositoryInterface,
	eventStoreRepository eventInterfaces.EventStoreRepositoryInterface,
	categoryQueryRepository categoryInterfaces.CategoryQueryRepositoryInterface,
) *EventService {
	return &EventService{
		eventQueryRepository:    eventQueryRepository,
		eventStoreRepository:    eventStoreRepository,
		categoryQueryRepository: categoryQueryRepository,
	}
}

func (service *EventService) Pagination(ctx context.Context, dto *eventDtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity] {
	return service.eventQueryRepository.Pagination(ctx, dto)
}

func (service *EventService) Detail(ctx context.Context, id uuid.UUID) *entities.EventEntity {
	data := service.eventQueryRepository.FindOneById(ctx, id)
	if data == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_NOT_FOUND))
	}

	return data
}

func (service *EventService) Create(ctx context.Context, dto *eventDtos.EventCreateRequestDto, organizerUserId uuid.UUID) *entities.EventEntity {
	categoryExists := service.categoryQueryRepository.FindOneById(ctx, dto.CategoryId)
	if categoryExists == nil {
		panic(*exceptions.BadRequestException(eventConstants.EVENT_CATEGORY_NOT_FOUND))
	}

	newEvent := dto.ToEntity(organizerUserId)

	isExistsByTitle := service.eventQueryRepository.IsExistsByTitle(ctx, newEvent.Title)
	if isExistsByTitle {
		panic(*exceptions.UnprocessableEntityException(eventConstants.EVENT_TITLE_EXISTS))
	}

	isExistsBySlug := service.eventQueryRepository.IsExistsBySlug(ctx, newEvent.Slug)
	if isExistsBySlug {
		panic(*exceptions.UnprocessableEntityException(eventConstants.EVENT_SLUG_EXISTS))
	}

	return service.eventStoreRepository.Create(ctx, newEvent)
}

func (service *EventService) Update(ctx context.Context, dto *eventDtos.EventUpdateRequestDto, currentUserId uuid.UUID) *entities.EventEntity {
	existingEvent := service.eventQueryRepository.FindOneById(ctx, dto.Id)
	if existingEvent == nil {
		panic(*exceptions.NotFoundException(eventConstants.EVENT_NOT_FOUND))
	}

	if existingEvent.OrganizerUserId != currentUserId {
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

	isExistsBySlug := service.eventQueryRepository.IsExistsBySlugExcludeId(ctx, updateEvent.Slug, updateEvent.Id)
	if isExistsBySlug {
		panic(*exceptions.BadRequestException(eventConstants.EVENT_SLUG_EXISTS))
	}

	return service.eventStoreRepository.Update(ctx, updateEvent)
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
