package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRegistrationEntity struct {
	Id            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"                                                                                              json:"id"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:uq_event_registrations_user_ticket,priority:1;index:idx_event_registrations_user_status,priority:1"           json:"user_id"`
	EventTicketID uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:uq_event_registrations_user_ticket,priority:2;index:idx_event_registrations_ticket_status,priority:1"         json:"event_ticket_id"`
	Status        int            `gorm:"not null;index:idx_event_registrations_user_status,priority:2;index:idx_event_registrations_ticket_status,priority:2"                        json:"status"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"                                                                                                                               json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"                                                                                                                               json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index"                                                                                                                                        json:"deleted_at,omitempty"`
	CreatedBy     *uuid.UUID     `gorm:"type:uuid"                                                                                                                                    json:"created_by,omitempty"`
	UpdatedBy     *uuid.UUID     `gorm:"type:uuid"                                                                                                                                    json:"updated_by,omitempty"`
	DeletedBy     *uuid.UUID     `gorm:"type:uuid"                                                                                                                                    json:"deleted_by,omitempty"`

	// Relations
	User        UserEntity        `gorm:"foreignKey:UserID"        json:"user,omitempty"`
	EventTicket EventTicketEntity `gorm:"foreignKey:EventTicketID" json:"event_ticket,omitempty"`
}

func (EventRegistrationEntity) TableName() string { return "event_registrations" }
