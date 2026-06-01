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
	roleDtos "event-backend/app/role/dtos"
)

type RoleQueryRepository struct {
	db        *gorm.DB
	roleModel *gorm.DB
}

func NewRoleQueryRepository(db *gorm.DB) *RoleQueryRepository {
	return &RoleQueryRepository{
		db:        db,
		roleModel: db.Model(&entities.RoleEntity{}),
	}
}

func (repo *RoleQueryRepository) Pagination(ctx context.Context, dto *roleDtos.RoleQueryRequestDto) *infradtos.PaginationResultDto[entities.RoleEntity] {
	query := repo.roleModel.WithContext(ctx)
	var results []*entities.RoleEntity
	var total int64

	query = repo.querySearch(query, dto)
	query = repo.querySort(query, dto)

	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count roles:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).Scopes(utils.Paginate(&dto.PaginationQueryRequestDto)).Find(&results).Error
	if err != nil {
		log.Println("Error find roles:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &infradtos.PaginationResultDto[entities.RoleEntity]{
		Data:  results,
		Count: total,
	}
}

func (repo *RoleQueryRepository) FindOneByName(ctx context.Context, name string) *entities.RoleEntity {
	query := repo.roleModel.WithContext(ctx)
	var result entities.RoleEntity

	err := query.Where("name = ?", name).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find role by name:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (repo *RoleQueryRepository) FindOneById(ctx context.Context, id uuid.UUID) *entities.RoleEntity {
	query := repo.roleModel.WithContext(ctx)
	var result entities.RoleEntity

	err := query.Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find role by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (repo *RoleQueryRepository) IsExistsById(ctx context.Context, id uuid.UUID) bool {
	query := repo.roleModel.WithContext(ctx)
	var exists bool
	err := query.Select("1").Where("id = ?", id).Scan(&exists).Error
	if err != nil {
		log.Println("Error check role exist by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}
	return exists
}

func (repo *RoleQueryRepository) IsExistsByName(ctx context.Context, name string) bool {
	query := repo.roleModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("name = ?", name).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check role exists by name:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (repo *RoleQueryRepository) IsExistsByNameExcludeId(ctx context.Context, name string, excludeID uuid.UUID) bool {
	query := repo.roleModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("name = ? AND id != ?", name, excludeID).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check role exists by name exclude id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (repo *RoleQueryRepository) querySearch(db *gorm.DB, dto *roleDtos.RoleQueryRequestDto) *gorm.DB {
	if dto.Search != "" {
		db = db.Where("name ILIKE ?", "%"+dto.Search+"%")
	}
	return db
}

func (repo *RoleQueryRepository) querySort(db *gorm.DB, dto *roleDtos.RoleQueryRequestDto) *gorm.DB {
	sortableColumns := []string{
		"name",
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
