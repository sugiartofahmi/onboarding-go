package services

import (
	"context"
	"time"

	categoryConstants "event-backend/app/category/constants"
	categoryInterfaces "event-backend/app/category/interfaces"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/exceptions"
	categoryDtos "event-backend/presentation/http/category/dtos"

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

func (service *CategoryService) Pagination(ctx context.Context, dto *categoryDtos.CategoryQueryRequestDto) *infradtos.PaginationResultDto[entities.CategoryEntity] {
	return service.categoryQueryRepository.Pagination(ctx, dto)
}

func (service *CategoryService) Detail(ctx context.Context, id uuid.UUID) *entities.CategoryEntity {
	data := service.categoryQueryRepository.FindOneById(ctx, id)
	if data == nil {
		panic(*exceptions.NotFoundException(categoryConstants.CATEGORY_NOT_FOUND))
	}

	return data
}

func (service *CategoryService) Create(ctx context.Context, dto *categoryDtos.CategoryCreateRequestDto) *entities.CategoryEntity {
	newCategory := dto.ToEntity()
	isExistsByName := service.categoryQueryRepository.IsExistsByName(ctx, newCategory.Name)

	if isExistsByName {
		panic(*exceptions.UnprocessableEntityException(categoryConstants.CATEGORY_NAME_EXISTS))
	}

	isExistsBySlug := service.categoryQueryRepository.IsExistsBySlug(ctx, newCategory.Slug)
	if isExistsBySlug {
		panic(*exceptions.UnprocessableEntityException(categoryConstants.CATEGORY_SLUG_EXISTS))
	}

	return service.categoryStoreRepository.Create(ctx, newCategory)
}

func (service *CategoryService) Update(ctx context.Context, dto *categoryDtos.CategoryUpdateRequestDto) *entities.CategoryEntity {
	existingCategory := service.categoryQueryRepository.FindOneById(ctx, dto.Id)
	if existingCategory == nil {
		panic(*exceptions.NotFoundException(categoryConstants.CATEGORY_NOT_FOUND))
	}

	updateCategory := dto.ToEntity(existingCategory)

	isExistsByName := service.categoryQueryRepository.IsExistsByNameExcludeId(ctx, updateCategory.Name, updateCategory.Id)
	if isExistsByName {
		panic(*exceptions.BadRequestException(categoryConstants.CATEGORY_NAME_EXISTS))
	}

	isExistsBySlug := service.categoryQueryRepository.IsExistsBySlugExcludeId(ctx, updateCategory.Slug, updateCategory.Id)
	if isExistsBySlug {
		panic(*exceptions.BadRequestException(categoryConstants.CATEGORY_SLUG_EXISTS))
	}

	return service.categoryStoreRepository.Update(ctx, updateCategory)
}

func (service *CategoryService) SoftDelete(ctx context.Context, id uuid.UUID) {
	category := service.categoryQueryRepository.FindOneById(ctx, id)
	if category == nil {
		panic(*exceptions.NotFoundException(categoryConstants.CATEGORY_NOT_FOUND))
	}

	now := time.Now()
	category.DeletedAt = gorm.DeletedAt{
		Time:  now,
		Valid: true,
	}

	service.categoryStoreRepository.Update(ctx, category)
}
