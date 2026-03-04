package repositories

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"event-backend/entities"
	"event-backend/infrastructure/exceptions"
)

type EventTicketQueryRepository struct {
	db               *gorm.DB
	eventTicketModel *gorm.DB
}

func NewEventTicketQueryRepository(db *gorm.DB) *EventTicketQueryRepository {
	return &EventTicketQueryRepository{
		db:               db,
		eventTicketModel: db.Model(&entities.EventTicketEntity{}),
	}
}

func (repo *EventTicketQueryRepository) WithTransaction(tx *gorm.DB) *EventTicketQueryRepository {
	return &EventTicketQueryRepository{
		db:               tx,
		eventTicketModel: tx.Model(&entities.EventTicketEntity{}),
	}
}

func (repo *EventTicketQueryRepository) FindManyByEventId(ctx context.Context, eventId uuid.UUID) []entities.EventTicketEntity {
	query := repo.eventTicketModel.WithContext(ctx)
	var results []entities.EventTicketEntity

	err := query.Where("event_id = ?", eventId).Find(&results).Error
	if err != nil {
		log.Println("Error find event tickets by event id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return results
}

func (repo *EventTicketQueryRepository) FindOneById(ctx context.Context, id uuid.UUID) *entities.EventTicketEntity {
	query := repo.eventTicketModel.WithContext(ctx)
	var result entities.EventTicketEntity

	err := query.Where("id = ?", id).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find event ticket by id:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

func (repo *EventTicketQueryRepository) FindOneByEventIdAndType(ctx context.Context, eventId uuid.UUID, ticketType int) *entities.EventTicketEntity {
	query := repo.eventTicketModel.WithContext(ctx)
	var result entities.EventTicketEntity

	err := query.Where("event_id = ? AND type = ?", eventId, ticketType).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		log.Println("Error find event ticket by event id and type:", err)
		panic(*exceptions.ServerErrorException(err))
	}

	return &result
}

