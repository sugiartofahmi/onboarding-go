package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"event-backend/app/event/interfaces"
	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type EventTicketStoreRepository struct {
	db               *gorm.DB
	eventTicketModel *gorm.DB
}

func NewEventTicketStoreRepository(db *gorm.DB) *EventTicketStoreRepository {
	return &EventTicketStoreRepository{
		db:               db,
		eventTicketModel: db.Model(&entities.EventTicketEntity{}),
	}
}

func (repo *EventTicketStoreRepository) WithTransaction(tx *gorm.DB) interfaces.EventTicketStoreRepositoryInterface {
	return &EventTicketStoreRepository{
		db:               tx,
		eventTicketModel: tx.Model(&entities.EventTicketEntity{}),
	}
}

func (repo *EventTicketStoreRepository) Create(ctx context.Context, entity *entities.EventTicketEntity) *entities.EventTicketEntity {
	query := repo.eventTicketModel.WithContext(ctx)

	err := query.Create(entity).Error
	if err != nil {
		log.Println("Error create event ticket:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (repo *EventTicketStoreRepository) Update(ctx context.Context, entity *entities.EventTicketEntity) *entities.EventTicketEntity {
	query := repo.eventTicketModel.WithContext(ctx)

	err := query.Where("id = ?", entity.Id).Updates(entity).Error
	if err != nil {
		log.Println("Error update event ticket:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entity
}

func (repo *EventTicketStoreRepository) Delete(ctx context.Context, id uuid.UUID) bool {
	query := repo.eventTicketModel.WithContext(ctx)

	result := query.Where("id = ?", id).Delete(&entities.EventTicketEntity{})
	if result.Error != nil {
		log.Println("Error delete event ticket:", result.Error)
		panic(*exceptions.ServerErrorException(result.Error))
	}

	return result.RowsAffected > 0
}

func (repo *EventTicketStoreRepository) BulkCreate(ctx context.Context, entities *[]entities.EventTicketEntity) *[]entities.EventTicketEntity {
	maxBatchSize := 20
	query := repo.eventTicketModel.WithContext(ctx)

	err := query.CreateInBatches(*entities, maxBatchSize).Error
	if err != nil {
		log.Println("Error bulk create event ticket:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return entities
}

func (repo *EventTicketStoreRepository) BulkDelete(ctx context.Context, entities *[]entities.EventTicketEntity) {
	query := repo.eventTicketModel.WithContext(ctx)

	err := query.Delete(entities).Error
	if err != nil {
		log.Println("Error bulk delete event ticket:", err)
		panic(*exceptions.ServerErrorException(err))
	}
}
