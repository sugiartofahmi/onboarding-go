package services

import (
	"context"
	"event-backend/app/auth/interfaces"
	roleConstants "event-backend/app/role/constants"
	roleInterfaces "event-backend/app/role/interfaces"
	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
	authdtos "event-backend/presentation/http/auth/dtos"
)

type AuthService struct {
	authQueryRepository interfaces.AuthQueryRepositoryInterface
	authStoreRepository interfaces.AuthStoreRepositoryInterface
	roleQueryRepository roleInterfaces.RoleQueryRepositoryInterface
}

func NewAuthService(
	authQueryRepository interfaces.AuthQueryRepositoryInterface,
	authStoreRepository interfaces.AuthStoreRepositoryInterface,
	roleQueryRepository roleInterfaces.RoleQueryRepositoryInterface,
) *AuthService {
	return &AuthService{
		authQueryRepository: authQueryRepository,
		authStoreRepository: authStoreRepository,
		roleQueryRepository: roleQueryRepository,
	}
}

func (service *AuthService) Login(ctx context.Context, request *authdtos.AuthLoginRequestDto) *authdtos.AuthLoginResponseDto {
	user := service.authQueryRepository.FindOneByEmailWithRole(ctx, request.Email)
	if user == nil {
		panic(*exceptions.UnauthenticatedException("Credential not valid"))
	}

	isPasswordValid := utils.ComparePassword(user.Password, request.Password)
	if !isPasswordValid {
		panic(*exceptions.UnauthenticatedException("Credential not valid"))
	}

	token, expiresAt := utils.GenerateToken(user)

	return &authdtos.AuthLoginResponseDto{
		Token:     token,
		TokenType: "Bearer",
		ExpiresAt: expiresAt,
	}
}

func (service *AuthService) Register(ctx context.Context, request *authdtos.AuthRegisterRequestDto) *authdtos.AuthRegisterResponseDto {
	isEmailExists := service.authQueryRepository.IsExistsByEmail(ctx, request.Email)
	if isEmailExists {
		panic(*exceptions.UnprocessableEntityException("Email already exists"))
	}

	role := service.roleQueryRepository.FindOneByName(ctx, roleConstants.ATTENDEE)
	if role == nil {
		panic(*exceptions.UnprocessableEntityException("Role not found"))
	}

	user := request.ToEntity()
	user.RoleId = role.Id
	result := service.authStoreRepository.Create(ctx, user)

	return &authdtos.AuthRegisterResponseDto{
		Id:        result.Id,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	}
}
