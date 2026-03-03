package services

import (
	"context"
	"time"

	userConstants "event-backend/app/user/constants"
	userInterfaces "event-backend/app/user/interfaces"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/exceptions"
	userdtos "event-backend/presentation/http/user/dtos"

	"github.com/google/uuid"
	"gorm.io/gorm"

	roleInterfaces "event-backend/app/role/interfaces"
)

type UserService struct {
	userQueryRepository userInterfaces.UserQueryRepositoryInterface
	userStoreRepository userInterfaces.UserStoreRepositoryInterface
	roleQueryRepository roleInterfaces.RoleQueryRepositoryInterface
}

func NewUserService(
	userQueryRepository userInterfaces.UserQueryRepositoryInterface,
	userStoreRepository userInterfaces.UserStoreRepositoryInterface,
	roleQueryRepository roleInterfaces.RoleQueryRepositoryInterface,
) *UserService {
	return &UserService{
		userQueryRepository: userQueryRepository,
		userStoreRepository: userStoreRepository,
		roleQueryRepository: roleQueryRepository,
	}
}


func (service *UserService) Pagination(ctx context.Context, dto *userdtos.UserQueryRequestDTO) *infradtos.PaginationResultDto[entities.UserEntity] {
	return service.userQueryRepository.Pagination(ctx, dto)
}

func (service *UserService) Detail(ctx context.Context, id uuid.UUID) *entities.UserEntity {
	data := service.userQueryRepository.FindOneById(ctx, id)
	if data == nil {
		panic(*exceptions.NotFoundException(userConstants.USER_NOT_FOUND))
	}

	return data
}

func (service *UserService) Create(ctx context.Context, dto *userdtos.UserCreateRequestDTO) *entities.UserEntity {
	isEmailExists := service.userQueryRepository.IsExistsByEmail(ctx, dto.Email)
	if isEmailExists {
		panic(*exceptions.BadRequestException(userConstants.USER_EMAIL_EXISTS))
	}

	isRoleExists := service.roleQueryRepository.IsExistsById(ctx, dto.RoleId)
	if !isRoleExists {
		panic(*exceptions.NotFoundException(userConstants.USER_ROLE_NOT_FOUND))
	}

	newUser := dto.ToEntity()
	return service.userStoreRepository.Create(ctx, newUser)
}

func (service *UserService) Update(ctx context.Context, dto *userdtos.UserUpdateRequestDTO) *entities.UserEntity {
	existingUser := service.userQueryRepository.FindOneById(ctx, dto.Id)
	if existingUser == nil {
		panic(*exceptions.NotFoundException(userConstants.USER_NOT_FOUND))
	}

	updateUser := dto.ToEntity(existingUser)

	isExistsByEmail := service.userQueryRepository.IsExistsByEmailExcludeId(ctx, updateUser.Email, updateUser.Id)
	if isExistsByEmail {
		panic(*exceptions.BadRequestException(userConstants.USER_EMAIL_EXISTS))
	}

	if dto.RoleId != nil {
		isRoleExists := service.roleQueryRepository.IsExistsById(ctx, *dto.RoleId)
		if !isRoleExists {
			panic(*exceptions.NotFoundException(userConstants.USER_ROLE_NOT_FOUND))
		}
	}

	return service.userStoreRepository.Update(ctx, updateUser)
}

func (service *UserService) Delete(ctx context.Context, id uuid.UUID) {
	existingUser := service.userQueryRepository.FindOneById(ctx, id)
	if existingUser == nil {
		panic(*exceptions.NotFoundException(userConstants.USER_NOT_FOUND))
	}

	now := time.Now()
	existingUser.DeletedAt = gorm.DeletedAt{
		Time:  now,
		Valid: true,
	}

	service.userStoreRepository.Update(ctx, existingUser)
}
