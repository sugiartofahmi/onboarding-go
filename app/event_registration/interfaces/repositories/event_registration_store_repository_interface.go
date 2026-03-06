package interfaces

import (
	"context"

	"gorm.io/gorm"

	"event-backend/entities"
)

type EventRegistrationStoreRepositoryInterface interface {
	Create(ctx context.Context, entity *entities.EventRegistrationEntity) *entities.EventRegistrationEntity
	Update(ctx context.Context, entity *entities.EventRegistrationEntity) *entities.EventRegistrationEntity
	WithTransaction(tx *gorm.DB) EventRegistrationStoreRepositoryInterface
}
