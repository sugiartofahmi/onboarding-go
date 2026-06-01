package services

import (
	"context"
	"time"

	"gorm.io/gorm"

	"event-backend/app/health/interfaces"
	"event-backend/infrastructure/config"
	redisInterfaces "event-backend/infrastructure/redis/interfaces"
	healthDtos "event-backend/presentation/http/health/dtos"
)

type HealthService struct {
	db         *gorm.DB
	redisCache redisInterfaces.RedisCacheInterface
}

func NewHealthService(db *gorm.DB, redisCache redisInterfaces.RedisCacheInterface) interfaces.HealthServiceInterface {
	return &HealthService{
		db:         db,
		redisCache: redisCache,
	}
}

func (s *HealthService) Check() healthDtos.HealthResponseDto {
	databaseStatus := "connected"
	redisStatus := "connected"

	sqlDB, err := s.db.DB()
	if err != nil || sqlDB.Ping() != nil {
		databaseStatus = "disconnected"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := s.redisCache.Set(ctx, "health_check", "1", 1*time.Second); err != nil {
		redisStatus = "disconnected"
	}

	return healthDtos.HealthResponseDto{
		App:         config.AppName,
		Environment: config.AppEnv,
		Database:    databaseStatus,
		Redis:       redisStatus,
	}
}
