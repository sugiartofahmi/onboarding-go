package repositories

import (
	"context"
	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
