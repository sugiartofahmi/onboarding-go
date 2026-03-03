package repositories

import (
	"context"
	"log"

	"gorm.io/gorm"

	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type UserStoreRepository struct {
	userModel *gorm.DB
}

func NewUserStoreRepository(db *gorm.DB) *UserStoreRepository {
	return &UserStoreRepository{
		userModel: db.Model(&entities.UserEntity{}),
	}
}

func (user *UserStoreRepository) Create(ctx context.Context, entity *entities.UserEntity) *entities.UserEntity {
	query := user.userModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create user:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (user *UserStoreRepository) Update(ctx context.Context, entity *entities.UserEntity) *entities.UserEntity {
	query := user.userModel.WithContext(ctx)

	err := query.Where("id = ?", entity.Id).Updates(entity).Error
	if err != nil {
		log.Println("Error update user:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}
