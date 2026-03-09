package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/enums"
	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
	eventdtos "event-backend/presentation/http/event/dtos"
)

type EventQueryRepository struct {
	db         *gorm.DB
	eventModel *gorm.DB
}

func NewEventQueryRepository(db *gorm.DB) *EventQueryRepository {
	return &EventQueryRepository{
		db:         db,
		eventModel: db.Model(&entities.EventEntity{}),
	}
}

func (repo *EventQueryRepository) Pagination(ctx context.Context, dto *eventdtos.EventQueryRequestDto) *infradtos.PaginationResultDto[entities.EventEntity] {
	query := repo.eventModel.WithContext(ctx)
	var results []*entities.EventEntity
	var total int64
	var eventPaginationColumns = []string{
	    "id",
	    "organizer_user_id",
	    "category_id",
	    "title",
	    "slug",
	    "description",
	    "location",
	    "start_date",
	    "end_date",
	    "status",
	    "created_at",
	    "updated_at",
	}

	query = query.Select(eventPaginationColumns)
	query = repo.querySearch(query, dto)
	query = repo.querySort(query, dto)
	query = repo.queryFilter(query, dto)

	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count events:", err)
		panic(*exceptions.ServerErrorException(err))
	}
	paginate := utils.Paginate(&dto.PaginationQueryRequestDto)
	err = query.Session(&gorm.Session{}).Scopes(paginate).Find(&results).Error
	if err != nil {
		log.Println("Error find events:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &infradtos.PaginationResultDto[entities.EventEntity]{
		Data:  results,
		Count: total,
	}
}

func (repo *EventQueryRepository) FindOneById(ctx context.Context, id uuid.UUID) *entities.EventEntity {
	query := repo.eventModel.WithContext(ctx)
	var result entities.EventEntity

	err := query.Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find event by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (repo *EventQueryRepository) FindOneBySlug(ctx context.Context, slug string) *entities.EventEntity {
	query := repo.eventModel.WithContext(ctx)
	var result entities.EventEntity

	err := query.Where("slug = ?", slug).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find event by slug:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (repo *EventQueryRepository) FindOneByIdWithTickets(ctx context.Context, id uuid.UUID) *entities.EventEntity {
	query := repo.eventModel.WithContext(ctx)
	var result entities.EventEntity

	err := query.Preload("Tickets").Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find event by id with tickets:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (repo *EventQueryRepository) IsExistsByTitle(ctx context.Context, title string) bool {
	query := repo.eventModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("title = ?", title).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check event exists by title:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (repo *EventQueryRepository) IsExistsByTitleExcludeId(ctx context.Context, title string, excludeID uuid.UUID) bool {
	query := repo.eventModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("title = ? AND id != ?", title, excludeID).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check event exists by title exclude id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (repo *EventQueryRepository) IsExistsBySlug(ctx context.Context, slug string) bool {
	query := repo.eventModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("slug = ?", slug).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check event exists by slug:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (repo *EventQueryRepository) IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool {
	query := repo.eventModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("slug = ? AND id != ?", slug, excludeID).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check event exists by slug exclude id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (repo *EventQueryRepository) IsExistsByCategoryId(ctx context.Context, categoryId uuid.UUID) bool {
	query := repo.eventModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("category_id = ?", categoryId).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check event exists by category id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (repo *EventQueryRepository) querySearch(db *gorm.DB, dto *eventdtos.EventQueryRequestDto) *gorm.DB {
	if dto.Search != "" {
		db = db.Where("title ILIKE ?", "%"+dto.Search+"%")
	}
	return db
}

func (repo *EventQueryRepository) querySort(db *gorm.DB, dto *eventdtos.EventQueryRequestDto) *gorm.DB {
	sortableColumns := []string{
		"title",
		"start_date",
		"created_at",
		"updated_at",
	}

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

	return db.Order(sortBy + " " + order)
}

func (repo *EventQueryRepository) queryFilter(db *gorm.DB, dto *eventdtos.EventQueryRequestDto) *gorm.DB {
	if dto.CategoryId != nil {
		db = db.Where("category_id = ?", *dto.CategoryId)
	}

	if dto.Status != nil {
		db = db.Where("status = ?", *dto.Status)
	}

	// Filter by date range (start_date and end_date)
	if dto.StartDate != nil && dto.EndDate != nil {
		db = db.Where("start_date >= ? AND end_date <= ?", *dto.StartDate, *dto.EndDate)
	}

	return db
}
