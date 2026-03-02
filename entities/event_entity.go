package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventEntity struct {
	Id              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"                                                                                                json:"id"`
	OrganizerUserID uuid.UUID      `gorm:"type:uuid;not null;index:idx_events_organizer_status,priority:1"                                                                               json:"organizer_user_id"`
	CategoryID      uuid.UUID      `gorm:"type:uuid;not null;index:idx_events_category_status,priority:1"                                                                                json:"category_id"`
	Title           string         `gorm:"size:255;not null"                                                                                                                              json:"title"`
	Slug            string         `gorm:"size:255;uniqueIndex;not null"                                                                                                                  json:"slug"`
	Description     *string        `gorm:"type:text"                                                                                                                                      json:"description,omitempty"`
	Location        *string        `gorm:"size:255"                                                                                                                                       json:"location,omitempty"`
	StartDate       time.Time      `gorm:"not null;index:idx_events_status_start_date,priority:2"                                                                                         json:"start_date"`
	EndDate         time.Time      `gorm:"not null"                                                                                                                                       json:"end_date"`
	Status          int            `gorm:"not null;index:idx_events_status_start_date,priority:1;index:idx_events_organizer_status,priority:2;index:idx_events_category_status,priority:2" json:"status"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"                                                                                                                                 json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"                                                                                                                                 json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"                                                                                                                                          json:"deleted_at,omitempty"`
	CreatedBy       *uuid.UUID     `gorm:"type:uuid"                                                                                                                                      json:"created_by,omitempty"`
	UpdatedBy       *uuid.UUID     `gorm:"type:uuid"                                                                                                                                      json:"updated_by,omitempty"`
	DeletedBy       *uuid.UUID     `gorm:"type:uuid"                                                                                                                                      json:"deleted_by,omitempty"`

	// Relations
	Organizer UserEntity          `gorm:"foreignKey:OrganizerUserID" json:"organizer,omitempty"`
	Category  CategoryEntity      `gorm:"foreignKey:CategoryID"      json:"category,omitempty"`
	Tickets   []EventTicketEntity `gorm:"foreignKey:EventID"         json:"tickets,omitempty"`
}

func (EventEntity) TableName() string { return "events" }
