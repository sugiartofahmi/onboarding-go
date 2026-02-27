package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventTicketEntity struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"                                                        json:"id"`
	EventID         uuid.UUID      `gorm:"type:uuid;not null;index:idx_event_tickets_event_id;uniqueIndex:uq_event_tickets_event_type,priority:1" json:"event_id"`
	Type            int            `gorm:"not null;uniqueIndex:uq_event_tickets_event_type,priority:2"                                            json:"type"`
	Price           float64        `gorm:"type:decimal(12,2);not null"                                                                            json:"price"`
	Quota           int            `gorm:"not null"                                                                                               json:"quota"`
	RegisteredCount int            `gorm:"not null;default:0"                                                                                     json:"registered_count"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"                                                                                         json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"                                                                                         json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"                                                                                                  json:"deleted_at,omitempty"`
	CreatedBy       *uuid.UUID     `gorm:"type:uuid"                                                                                              json:"created_by,omitempty"`
	UpdatedBy       *uuid.UUID     `gorm:"type:uuid"                                                                                              json:"updated_by,omitempty"`
	DeletedBy       *uuid.UUID     `gorm:"type:uuid"                                                                                              json:"deleted_by,omitempty"`

	// Relations
	Event         EventEntity               `gorm:"foreignKey:EventID"       json:"event,omitempty"`
	Registrations []EventRegistrationEntity `gorm:"foreignKey:EventTicketID" json:"registrations,omitempty"`
}

func (EventTicketEntity) TableName() string { return "event_tickets" }
