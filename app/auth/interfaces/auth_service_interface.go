package interfaces

import (
	"context"

	authdtos "event-backend/presentation/http/auth/dtos"
)

type AuthServiceInterface interface {
	Login(ctx context.Context, dto *authdtos.AuthLoginRequestDto) *authdtos.AuthLoginResponseDto
	Register(ctx context.Context, dto *authdtos.AuthRegisterRequestDto) *authdtos.AuthRegisterResponseDto
}