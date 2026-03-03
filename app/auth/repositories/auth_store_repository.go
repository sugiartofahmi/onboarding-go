package repositories

import (
	"context"
	"log"

	"gorm.io/gorm"

	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type AuthStoreRepository struct {
	db        *gorm.DB
	userModel *gorm.DB
}

func NewAuthStoreRepository(db *gorm.DB) *AuthStoreRepository {
	return &AuthStoreRepository{
		db:        db,
		userModel: db.Model(&entities.UserEntity{}),
	}
}

func (auth *AuthStoreRepository) Create(ctx context.Context, user *entities.UserEntity) *entities.UserEntity {
	query := auth.userModel.WithContext(ctx)

	err := query.Create(user).Error
	if err != nil {
		log.Println("Error create user:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return user
}
