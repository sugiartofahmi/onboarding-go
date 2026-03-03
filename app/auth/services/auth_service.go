package services

import (
	"context"
	"event-backend/app/auth/interfaces"
	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
	authdtos "event-backend/presentation/http/auth/dtos"
)

type AuthService struct {
	authQueryRepository interfaces.AuthQueryRepositoryInterface
}

func NewAuthService(
	authQueryRepository interfaces.AuthQueryRepositoryInterface,
) *AuthService {
	return &AuthService{
		authQueryRepository: authQueryRepository,
	}
}

func (service *AuthService) Login(ctx context.Context, email, password string) *authdtos.AuthLoginResponseDto {
	user := service.authQueryRepository.FindOneByEmailWithRole(ctx, email)
	if user == nil {
		panic(*exceptions.UnauthenticatedException("Credential not valid"))
	}

	isPasswordValid := utils.ComparePassword(user.Password, password)
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
