package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type RoleStoreRepository struct {
	roleModel *gorm.DB
}

func NewRoleStoreRepository(db *gorm.DB) *RoleStoreRepository {
	return &RoleStoreRepository{
		roleModel: db.Model(&entities.RoleEntity{}),
	}
}

func (repo *RoleStoreRepository) Create(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity {
	query := repo.roleModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create role:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (repo *RoleStoreRepository) Update(ctx context.Context, entity *entities.RoleEntity) *entities.RoleEntity {
	query := repo.roleModel.WithContext(ctx)

	err := query.Where("id = ?", entity.Id).Updates(entity).Error
	if err != nil {
		log.Println("Error update role:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (repo *RoleStoreRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := repo.roleModel.WithContext(ctx)
	entity := &entities.RoleEntity{Id: id}

	err := query.Delete(entity).Error
	if err != nil {
		log.Println("Error delete role:", err)
		return err
	}

	return nil
}
