package repositories

import (
	"context"
	"log"

	"gorm.io/gorm"

	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type CategoryStoreRepository struct {
	categoryModel *gorm.DB
}

func NewCategoryStoreRepository(db *gorm.DB) *CategoryStoreRepository {
	return &CategoryStoreRepository{
		categoryModel: db.Model(&entities.CategoryEntity{}),
	}
}

func (category *CategoryStoreRepository) Create(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity {
	query := category.categoryModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create category:", err)
		panic(exceptions.ServerErrorException(err))
	}

	return entity
}

func (category *CategoryStoreRepository) Update(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity {
	query := category.categoryModel.WithContext(ctx)

	err := query.Where("id = ?", entity.Id).Updates(entity).Error
	if err != nil {
		log.Println("Error update category:", err)
		panic(exceptions.ServerErrorException(err))
	}

	return entity
}
