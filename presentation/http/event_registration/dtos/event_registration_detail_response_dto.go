package dtos

import (
	"time"

	eventenums "event-backend/app/event/enums"
	eventregistrationenums "event-backend/app/event_registration/enums"
	"event-backend/entities"

	"github.com/google/uuid"
)

type EventRegistrationDetailResponseDto struct {
	Id            uuid.UUID                 `json:"id"`
	UserId        uuid.UUID                 `json:"user_id"`
	EventTicketId uuid.UUID                 `json:"event_ticket_id"`
	EventId       uuid.UUID                 `json:"event_id"`
	Status        int                       `json:"status"`
	StatusLabel   string                    `json:"status_label"`
	User          UserResponseDto           `json:"user,omitempty"`
	EventTicket   EventTicketResponseDto    `json:"event_ticket,omitempty"`
	Event         EventRegistrationEventDto `json:"event,omitempty"`
	CreatedAt     time.Time                 `json:"created_at"`
	UpdatedAt     time.Time                 `json:"updated_at"`
}

type UserResponseDto struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type EventTicketResponseDto struct {
	Id              uuid.UUID `json:"id"`
	EventId         uuid.UUID `json:"event_id"`
	Type            int       `json:"type"`
	TypeLabel       string    `json:"type_label"`
	Price           float64   `json:"price"`
	Quota           int       `json:"quota"`
	RegisteredCount int       `json:"registered_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type EventRegistrationEventDto struct {
	Id     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Slug   string    `json:"slug"`
	Status int       `json:"status"`
}

func EventRegistrationDetailResponseDtoFromEntity(entity entities.EventRegistrationEntity) EventRegistrationDetailResponseDto {
	return EventRegistrationDetailResponseDto{
		Id:            entity.Id,
		UserId:        entity.UserId,
		EventTicketId: entity.EventTicketId,
		EventId:       entity.EventId,
		Status:        entity.Status,
		StatusLabel:   eventregistrationenums.EventRegistrationStatusEnum(entity.Status).GetLabel(),
		User: UserResponseDto{
			Id:    entity.User.Id,
			Name:  entity.User.Name,
			Email: entity.User.Email,
		},
		EventTicket: EventTicketResponseDto{
			Id:              entity.EventTicket.Id,
			EventId:         entity.EventTicket.EventId,
			Type:            entity.EventTicket.Type,
			TypeLabel:       eventenums.EventTicketTypeEnum(entity.EventTicket.Type).GetLabel(),
			Price:           entity.EventTicket.Price,
			Quota:           entity.EventTicket.Quota,
			RegisteredCount: entity.EventTicket.RegisteredCount,
			CreatedAt:       entity.EventTicket.CreatedAt,
			UpdatedAt:       entity.EventTicket.UpdatedAt,
		},
		Event: EventRegistrationEventDto{
			Id:     entity.Event.Id,
			Title:  entity.Event.Title,
			Slug:   entity.Event.Slug,
			Status: entity.Event.Status,
		},
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
