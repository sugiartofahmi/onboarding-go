package interfaces

import healthDtos "event-backend/app/health/dtos"

type HealthServiceInterface interface {
	Check() healthDtos.HealthResponseDto
}
