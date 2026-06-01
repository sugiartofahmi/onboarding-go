package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
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
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (category *CategoryStoreRepository) Update(ctx context.Context, entity *entities.CategoryEntity) *entities.CategoryEntity {
	query := category.categoryModel.WithContext(ctx)

	err := query.Where("id = ?", entity.Id).Updates(entity).Error
	if err != nil {
		log.Println("Error update category:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (category *CategoryStoreRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := category.categoryModel.WithContext(ctx)
	entity := &entities.CategoryEntity{Id: id}

	err := query.Delete(entity).Error
	if err != nil {
		log.Println("Error delete category:", err)
		return err
	}

	return nil
}
