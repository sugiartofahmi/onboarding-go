package services

import (
	"context"
	categoryDtos "event-backend/app/category/dtos"
	categoryInterfaces "event-backend/app/category/interfaces"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/exceptions"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryService struct {
	categoryQueryRepository categoryInterfaces.CategoryQueryRepositoryInterface
	categoryStoreRepository categoryInterfaces.CategoryStoreRepositoryInterface
}

func NewCategoryService(
	categoryQueryRepository categoryInterfaces.CategoryQueryRepositoryInterface,
	categoryStoreRepository categoryInterfaces.CategoryStoreRepositoryInterface,
) *CategoryService {
	return &CategoryService{
		categoryQueryRepository: categoryQueryRepository,
		categoryStoreRepository: categoryStoreRepository,
	}
}

func (service *CategoryService) Pagination(ctx context.Context, dto *categoryDtos.CategoryQueryRequestDTO) *infradtos.PaginationResultDto[entities.CategoryEntity] {
	return service.categoryQueryRepository.Pagination(ctx, dto)
}

func (service *CategoryService) Detail(ctx context.Context, id uuid.UUID) *entities.CategoryEntity {
	data := service.categoryQueryRepository.FindOneById(ctx, id)
	if data == nil {
		panic(exceptions.NotFoundException("Category not found"))
	}

	return data
}

func (service *CategoryService) Create(ctx context.Context, dto *categoryDtos.CategoryCreateRequestDTO) *entities.CategoryEntity {
	newCategory := dto.ToEntity()
	isExistsByName := service.categoryQueryRepository.IsExistsByName(ctx, newCategory.Name)

	if isExistsByName {
		panic(exceptions.UnprocessableEntityException("Category name already exists"))
	}

	isExistsBySlug := service.categoryQueryRepository.IsExistsBySlug(ctx, newCategory.Slug)
	if isExistsBySlug {
		panic(exceptions.UnprocessableEntityException("Category slug already exists"))
	}

	return service.categoryStoreRepository.Create(ctx, newCategory)
}


func (service *CategoryService) Update(ctx context.Context, dto *categoryDtos.CategoryUpdateRequestDTO) *entities.CategoryEntity {
	existingCategory := service.categoryQueryRepository.FindOneById(ctx, dto.Id)
	if existingCategory == nil {
		panic(exceptions.NotFoundException("Category not found"))
	}

	updateCategory := dto.ToEntity(existingCategory)

	isExistsByName := service.categoryQueryRepository.IsExistsByNameExcludeId(ctx, updateCategory.Name, updateCategory.Id)
	if isExistsByName {
		panic(exceptions.BadRequestException("Category name already exists"))
	}

	isExistsBySlug := service.categoryQueryRepository.IsExistsBySlugExcludeId(ctx, updateCategory.Slug, updateCategory.Id)
	if isExistsBySlug {
		panic(exceptions.BadRequestException("Category slug already exists"))
	}

	return service.categoryStoreRepository.Update(ctx, updateCategory)
}

func (service *CategoryService) SoftDelete(ctx context.Context, id uuid.UUID) {
	category := service.categoryQueryRepository.FindOneById(ctx, id)
	if category == nil {
		panic(exceptions.NotFoundException("Category not found"))
	}

	now := time.Now()
	category.DeletedAt = gorm.DeletedAt{
        Time:  now,
        Valid: true,
    }

	service.categoryStoreRepository.Update(ctx, category)
}