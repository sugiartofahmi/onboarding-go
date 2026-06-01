package interfaces

import "event-backend/presentation/http/health/dtos"

type HealthServiceInterface interface {
	Check() dtos.HealthResponseDto
}
