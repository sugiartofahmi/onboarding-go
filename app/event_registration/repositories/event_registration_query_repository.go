package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	roleConstants "event-backend/app/role/constants"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/enums"
	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
	eventregistrationdtos "event-backend/presentation/http/event_registration/dtos"
)

type EventRegistrationQueryRepository struct {
	db            *gorm.DB
	eventRegModel *gorm.DB
}

func NewEventRegistrationQueryRepository(db *gorm.DB) *EventRegistrationQueryRepository {
	return &EventRegistrationQueryRepository{
		db:            db,
		eventRegModel: db.Model(&entities.EventRegistrationEntity{}),
	}
}

func (repo *EventRegistrationQueryRepository) Pagination(ctx context.Context, dto *eventregistrationdtos.EventRegistrationQueryRequestDto) *infradtos.PaginationResultDto[entities.EventRegistrationEntity] {
	query := repo.eventRegModel.WithContext(ctx)
	var results []*entities.EventRegistrationEntity
	var total int64

	query = repo.querySearch(query, dto)
	query = repo.querySort(query, dto)
	query = repo.queryFilter(query, dto)

	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count event registrations:", err)
		panic(*exceptions.ServerErrorException(err))
	}
	paginate := utils.Paginate(&dto.PaginationQueryRequestDto)
	err = query.Session(&gorm.Session{}).Scopes(paginate).Preload("User").Preload("Event").Preload("EventTicket").Find(&results).Error
	if err != nil {
		log.Println("Error find event registrations:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &infradtos.PaginationResultDto[entities.EventRegistrationEntity]{
		Data:  results,
		Count: total,
	}
}

func (repo *EventRegistrationQueryRepository) FindOneById(ctx context.Context, id uuid.UUID) *entities.EventRegistrationEntity {
	query := repo.eventRegModel.WithContext(ctx)
	var result entities.EventRegistrationEntity

	err := query.Preload("User").Preload("Event").Preload("EventTicket").Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find event registration by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (repo *EventRegistrationQueryRepository) querySearch(query *gorm.DB, dto *eventregistrationdtos.EventRegistrationQueryRequestDto) *gorm.DB {

	return query
}

func (repo *EventRegistrationQueryRepository) querySort(query *gorm.DB, dto *eventregistrationdtos.EventRegistrationQueryRequestDto) *gorm.DB {
	sortableColumns := []string{"status", "created_at", "updated_at"}

	sortBy := "created_at"
	if dto.SortBy != "" {
		isColumnAllowed := utils.Contains(sortableColumns, dto.SortBy)
		if isColumnAllowed {
			sortBy = dto.SortBy
		}
	}

	order := "DESC"
	if dto.Order == enums.SortOrderAsc {
		order = "ASC"
	}

	return query.Order(sortBy + " " + order)
}

func (repo *EventRegistrationQueryRepository) queryFilter(query *gorm.DB, dto *eventregistrationdtos.EventRegistrationQueryRequestDto) *gorm.DB {
	isRoleAdmin := dto.CurrentUserRoleName == roleConstants.ADMIN
	if !isRoleAdmin {
		query = query.Where("user_id = ?", dto.CurrentUserId)
	}

	if dto.EventId != "" {
		if eventId, err := uuid.Parse(dto.EventId); err == nil {
			query = query.Where("event_id = ?", eventId)
		}
	}

	if dto.Status != "" {
		query = query.Where("status = ?", dto.Status)
	}
	return query
}
