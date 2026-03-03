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
	userdtos "event-backend/presentation/http/user/dtos"
)

type UserQueryRepository struct {
	db        *gorm.DB
	userModel *gorm.DB
}

func NewUserQueryRepository(db *gorm.DB) *UserQueryRepository {
	return &UserQueryRepository{
		db:        db,
		userModel: db.Model(&entities.UserEntity{}),
	}
}

func (user *UserQueryRepository) Pagination(ctx context.Context, dto *userdtos.UserQueryRequestDto) *infradtos.PaginationResultDto[entities.UserEntity] {
	query := user.userModel.WithContext(ctx)
	var results []*entities.UserEntity
	var total int64

	query = user.QuerySearch(query, dto)
	query = user.QuerySort(query, dto)

	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count users:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).Preload("Role").Scopes(utils.Paginate(&dto.PaginationQueryRequestDto)).Find(&results).Error
	if err != nil {
		log.Println("Error find users:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &infradtos.PaginationResultDto[entities.UserEntity]{
		Data:  results,
		Count: total,
	}
}

func (user *UserQueryRepository) FindOneById(ctx context.Context, id uuid.UUID) *entities.UserEntity {
	query := user.userModel.WithContext(ctx)
	var result entities.UserEntity

	err := query.Preload("Role").Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find user by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (user *UserQueryRepository) FindOneByIdWithRole(ctx context.Context, id uuid.UUID) *entities.UserEntity {
	query := user.userModel.WithContext(ctx)
	var result entities.UserEntity

	err := query.Preload("Role").Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find user by id with role:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (user *UserQueryRepository) FindOneByEmail(ctx context.Context, email string) *entities.UserEntity {
	query := user.userModel.WithContext(ctx)
	var result entities.UserEntity

	err := query.Preload("Role").Where("email = ?", email).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find user by email:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (user *UserQueryRepository) IsExistsByEmail(ctx context.Context, email string) bool {
	query := user.userModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("email = ?", email).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check user exists by email:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (user *UserQueryRepository) IsExistsByEmailExcludeId(ctx context.Context, email string, excludeID uuid.UUID) bool {
	query := user.userModel.WithContext(ctx)
	var exists bool

	err := query.
		Select("1").
		Where("email = ? AND id != ?", email, excludeID).
		Limit(1).
		Scan(&exists).Error

	if err != nil {
		log.Println("Error check user exists by email exclude id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return exists
}

func (user *UserQueryRepository) QuerySearch(db *gorm.DB, dto *userdtos.UserQueryRequestDto) *gorm.DB {
	if dto.Search != "" {
		db = db.Where("name ILIKE ? OR email ILIKE ?", "%"+dto.Search+"%", "%"+dto.Search+"%")
	}
	return db
}

func (user *UserQueryRepository) QuerySort(db *gorm.DB, dto *userdtos.UserQueryRequestDto) *gorm.DB {
	allowedSortFields := map[string]bool{
		"name":       true,
		"email":      true,
		"is_active":  true,
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

func (user *UserQueryRepository) FindOneByRoleId(ctx context.Context, roleId uuid.UUID) *entities.RoleEntity {
	query := user.db.WithContext(ctx)
	var result entities.RoleEntity

	err := query.Where("id = ?", roleId).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find role by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}
