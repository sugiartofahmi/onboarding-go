package repositories

import (
	"context"
	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
	"log"

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
