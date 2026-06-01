package services

import (
	"context"
	authConstants "event-backend/app/auth/constants"
	"event-backend/app/auth/interfaces"
	roleConstants "event-backend/app/role/constants"
	roleInterfaces "event-backend/app/role/interfaces"
	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
	authDtos "event-backend/app/auth/dtos"
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

func (service *AuthService) Login(ctx context.Context, request *authDtos.AuthLoginRequestDto) *authDtos.AuthLoginResponseDto {
	user := service.authQueryRepository.FindOneByEmailWithRole(ctx, request.Email)
	if user == nil {
		panic(*exceptions.UnauthenticatedException(authConstants.AUTH_CREDENTIAL_NOT_VALID))
	}

	isPasswordValid := utils.ComparePassword(user.Password, request.Password)
	if !isPasswordValid {
		panic(*exceptions.UnauthenticatedException(authConstants.AUTH_CREDENTIAL_NOT_VALID))
	}

	token, expiresAt := utils.GenerateToken(user)

	return &authDtos.AuthLoginResponseDto{
		Token:     token,
		TokenType: "Bearer",
		ExpiresAt: expiresAt,
	}
}

func (service *AuthService) Register(ctx context.Context, request *authDtos.AuthRegisterRequestDto) *authDtos.AuthRegisterResponseDto {
	isEmailExists := service.authQueryRepository.IsExistsByEmail(ctx, request.Email)
	if isEmailExists {
		panic(*exceptions.UnprocessableEntityException(authConstants.AUTH_EMAIL_ALREADY_EXISTS))
	}

	role := service.roleQueryRepository.FindOneByName(ctx, roleConstants.ATTENDEE)
	if role == nil {
		panic(*exceptions.UnprocessableEntityException(authConstants.AUTH_ROLE_NOT_FOUND))
	}

	user := request.ToEntity()
	user.RoleId = role.Id
	result := service.authStoreRepository.Create(ctx, user)

	return &authDtos.AuthRegisterResponseDto{
		Id:        result.Id,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	}
}
