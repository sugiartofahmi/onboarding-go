package interfaces

import (
	"context"

	authDtos "event-backend/app/auth/dtos"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, dto *authDtos.AuthLoginRequestDto) *authDtos.AuthLoginResponseDto
	Register(ctx context.Context, dto *authDtos.AuthRegisterRequestDto) *authDtos.AuthRegisterResponseDto
}