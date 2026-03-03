package services

import (
	"context"
	"time"

	roleConstants "event-backend/app/role/constants"
	roleInterfaces "event-backend/app/role/interfaces"
	"event-backend/entities"
	infradtos "event-backend/infrastructure/dtos"
	"event-backend/infrastructure/exceptions"
	roleDtos "event-backend/presentation/http/role/dtos"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleService struct {
	roleQueryRepository roleInterfaces.RoleQueryRepositoryInterface
	roleStoreRepository roleInterfaces.RoleStoreRepositoryInterface
}

func NewRoleService(
	roleQueryRepository roleInterfaces.RoleQueryRepositoryInterface,
	roleStoreRepository roleInterfaces.RoleStoreRepositoryInterface,
) *RoleService {
	return &RoleService{
		roleQueryRepository: roleQueryRepository,
		roleStoreRepository: roleStoreRepository,
	}
}

func (service *RoleService) Pagination(ctx context.Context, dto *roleDtos.RoleQueryRequestDTO) *infradtos.PaginationResultDto[entities.RoleEntity] {
	return service.roleQueryRepository.Pagination(ctx, dto)
}

func (service *RoleService) Detail(ctx context.Context, id uuid.UUID) *entities.RoleEntity {
	data := service.roleQueryRepository.FindOneById(ctx, id)
	if data == nil {
		panic(*exceptions.NotFoundException(roleConstants.ROLE_NOT_FOUND))
	}

	return data
}

func (service *RoleService) Create(ctx context.Context, dto *roleDtos.RoleCreateRequestDTO) *entities.RoleEntity {
	newRole := dto.ToEntity()
	isExistsByName := service.roleQueryRepository.IsExistsByName(ctx, newRole.Name)

	if isExistsByName {
		panic(*exceptions.UnprocessableEntityException(roleConstants.ROLE_NAME_EXISTS))
	}

	return service.roleStoreRepository.Create(ctx, newRole)
}

func (service *RoleService) Update(ctx context.Context, dto *roleDtos.RoleUpdateRequestDTO) *entities.RoleEntity {
	existingRole := service.roleQueryRepository.FindOneById(ctx, dto.Id)
	if existingRole == nil {
		panic(*exceptions.NotFoundException(roleConstants.ROLE_NOT_FOUND))
	}

	updateRole := dto.ToEntity(existingRole)

	isExistsByName := service.roleQueryRepository.IsExistsByNameExcludeId(ctx, updateRole.Name, updateRole.Id)
	if isExistsByName {
		panic(*exceptions.BadRequestException(roleConstants.ROLE_NAME_EXISTS))
	}

	return service.roleStoreRepository.Update(ctx, updateRole)
}

func (service *RoleService) SoftDelete(ctx context.Context, id uuid.UUID) {
	role := service.roleQueryRepository.FindOneById(ctx, id)
	if role == nil {
		panic(*exceptions.NotFoundException(roleConstants.ROLE_NOT_FOUND))
	}

	now := time.Now()
	role.DeletedAt = gorm.DeletedAt{
		Time:  now,
		Valid: true,
	}

	service.roleStoreRepository.Update(ctx, role)
}
