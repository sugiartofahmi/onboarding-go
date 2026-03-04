package repositories

import (
	"context"
	"log"

	"gorm.io/gorm"

	"event-backend/app/event/interfaces"
	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type EventStoreRepository struct {
	db         *gorm.DB
	eventModel *gorm.DB
}

func NewEventStoreRepository(db *gorm.DB) *EventStoreRepository {
	return &EventStoreRepository{
		db:         db,
		eventModel: db.Model(&entities.EventEntity{}),
	}
}

func (repo *EventStoreRepository) WithTransaction(tx *gorm.DB) interfaces.EventStoreRepositoryInterface {
	return &EventStoreRepository{
		db:         tx,
		eventModel: tx.Model(&entities.EventEntity{}),
	}
}

func (repo *EventStoreRepository) Create(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity {
	query := repo.eventModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create event:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (repo *EventStoreRepository) Update(ctx context.Context, entity *entities.EventEntity) *entities.EventEntity {
	query := repo.eventModel.WithContext(ctx)

	err := query.Where("id = ?", entity.Id).Updates(entity).Error
	if err != nil {
		log.Println("Error update event:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}
