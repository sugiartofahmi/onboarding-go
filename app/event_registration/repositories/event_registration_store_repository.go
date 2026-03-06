package repositories

import (
	"context"
	"log"

	"gorm.io/gorm"

	eventregistrationInterfaces "event-backend/app/event_registration/interfaces/repositories"
	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type EventRegistrationStoreRepository struct {
	db            *gorm.DB
	eventRegModel *gorm.DB
}

func NewEventRegistrationStoreRepository(db *gorm.DB) *EventRegistrationStoreRepository {
	return &EventRegistrationStoreRepository{
		db:            db,
		eventRegModel: db.Model(&entities.EventRegistrationEntity{}),
	}
}

func (repo *EventRegistrationStoreRepository) WithTransaction(tx *gorm.DB) eventregistrationInterfaces.EventRegistrationStoreRepositoryInterface {
	return &EventRegistrationStoreRepository{
		db:            tx,
		eventRegModel: tx.Model(&entities.EventRegistrationEntity{}),
	}
}

func (repo *EventRegistrationStoreRepository) Create(ctx context.Context, entity *entities.EventRegistrationEntity) *entities.EventRegistrationEntity {
	query := repo.eventRegModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create event registration:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (repo *EventRegistrationStoreRepository) Update(ctx context.Context, entity *entities.EventRegistrationEntity) *entities.EventRegistrationEntity {
	query := repo.eventRegModel.WithContext(ctx)

	err := query.Where("id = ?", entity.Id).Updates(entity).Error
	if err != nil {
		log.Println("Error update event registration:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}