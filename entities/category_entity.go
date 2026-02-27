package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryEntity struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string         `gorm:"size:255;uniqueIndex;not null"                   json:"name"`
	Slug      string         `gorm:"size:255;uniqueIndex;not null"                   json:"slug"`
	CreatedAt time.Time      `gorm:"autoCreateTime"                                  json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"                                  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                           json:"deleted_at,omitempty"`
	CreatedBy *uuid.UUID     `gorm:"type:uuid"                                       json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID     `gorm:"type:uuid"                                       json:"updated_by,omitempty"`
	DeletedBy *uuid.UUID     `gorm:"type:uuid"                                       json:"deleted_by,omitempty"`

	// Relations
	Events []EventEntity `gorm:"foreignKey:CategoryID" json:"events,omitempty"`
}

func (CategoryEntity) TableName() string { return "categories" }
