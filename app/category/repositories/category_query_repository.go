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
	categorydtos "event-backend/presentation/http/category/dtos"
)

type CategoryQueryRepository struct {
	db            *gorm.DB
	categoryModel *gorm.DB
}

func NewCategoryQueryRepository(db *gorm.DB) *CategoryQueryRepository {
	return &CategoryQueryRepository{
		db:            db,
		categoryModel: db.Model(&entities.CategoryEntity{}),
	}
}

func (category *CategoryQueryRepository) Pagination(ctx context.Context, dto *categorydtos.CategoryQueryRequestDto) *infradtos.PaginationResultDto[entities.CategoryEntity] {
	query := category.categoryModel.WithContext(ctx)
	var results []*entities.CategoryEntity
	var total int64

	query = category.QuerySearch(query, dto)
	query = category.QuerySort(query, dto)

	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count categories:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).Scopes(utils.Paginate(&dto.PaginationQueryRequestDto)).Find(&results).Error
	if err != nil {
		log.Println("Error find categories:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &infradtos.PaginationResultDto[entities.CategoryEntity]{
		Data:  results,
		Count: total,
	}
}

func (category *CategoryQueryRepository) FindOneById(ctx context.Context, id uuid.UUID) *entities.CategoryEntity {
	query := category.categoryModel.WithContext(ctx)
	var result entities.CategoryEntity

	err := query.Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find category by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (category *CategoryQueryRepository) FindOneBySlug(ctx context.Context, slug string) *entities.CategoryEntity {
	query := category.categoryModel.WithContext(ctx)
	var result entities.CategoryEntity

	err := query.Where("slug = ?", slug).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find category by slug:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (category *CategoryQueryRepository) IsExistsByName(ctx context.Context, name string) bool {
	query := category.categoryModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("name = ?", name).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check category exists by name:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (category *CategoryQueryRepository) IsExistsByNameExcludeId(ctx context.Context, name string, excludeID uuid.UUID) bool {
	query := category.categoryModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("name = ? AND id != ?", name, excludeID).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check category exists by name exclude id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (category *CategoryQueryRepository) IsExistsBySlug(ctx context.Context, slug string) bool {
	query := category.categoryModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("slug = ?", slug).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check category exists by slug:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (category *CategoryQueryRepository) IsExistsBySlugExcludeId(ctx context.Context, slug string, excludeID uuid.UUID) bool {
	query := category.categoryModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("slug = ? AND id != ?", slug, excludeID).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check category exists by slug exclude id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

// Private helper methods

func (category *CategoryQueryRepository) QuerySearch(db *gorm.DB, dto *categorydtos.CategoryQueryRequestDto) *gorm.DB {
	if dto.Search != "" {
		db = db.Where("name ILIKE ?", "%"+dto.Search+"%")
	}
	return db
}

func (category *CategoryQueryRepository) QuerySort(db *gorm.DB, dto *categorydtos.CategoryQueryRequestDto) *gorm.DB {
	allowedSortFields := map[string]bool{
		"name":       true,
		"created_at": true,
		"updated_at": true,
	}

	sortBy := dto.SortBy
	if sortBy == "" || !allowedSortFields[sortBy] {
		sortBy = "created_at"
	}

	order := "DESC"
	if dto.Order == enums.SortOrderAsc {
		order = "ASC"
	}

	return db.Order(sortBy + " " + order)
}
