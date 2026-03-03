package repositories

import (
	"context"
	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
	"log"

	"gorm.io/gorm"
)

type AuthQueryRepository struct {
	db *gorm.DB
	userModel *gorm.DB
}

func NewAuthQueryRepository(db *gorm.DB) *AuthQueryRepository {
	return &AuthQueryRepository{
		db: db,
		userModel: db.Model(&entities.UserEntity{}),
	}
}

func (auth *AuthQueryRepository) FindOneByEmail(ctx context.Context, email string) *entities.UserEntity {
	query := auth.userModel.WithContext(ctx)
	var result entities.UserEntity

	err := query.Where("email = ?", email).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find user by email:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (auth *AuthQueryRepository) FindOneByEmailWithRole(ctx context.Context, email string) *entities.UserEntity {
	query := auth.userModel.WithContext(ctx)
	var result entities.UserEntity

	err := query.Preload("Role").Where("email = ?", email).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find user by email with role:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}