package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	Id        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"    json:"id"`
	RoleID    uuid.UUID      `gorm:"type:uuid;not null;index:idx_users_role_id"         json:"role_id"`
	Name      string         `gorm:"size:255;not null"                                  json:"name"`
	Email     string         `gorm:"size:255;uniqueIndex;not null"                      json:"email"`
	Password  string         `gorm:"size:255;not null"                                  json:"-"`
	IsActive  bool           `gorm:"not null;default:true"                              json:"is_active"`
	CreatedAt time.Time      `gorm:"autoCreateTime"                                     json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"                                     json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"                                              json:"deleted_at,omitempty"`
	CreatedBy *uuid.UUID     `gorm:"type:uuid"                                          json:"created_by,omitempty"`
	UpdatedBy *uuid.UUID     `gorm:"type:uuid"                                          json:"updated_by,omitempty"`
	DeletedBy *uuid.UUID     `gorm:"type:uuid"                                          json:"deleted_by,omitempty"`

	// Relations
	Role          RoleEntity                `gorm:"foreignKey:RoleID"          json:"role,omitempty"`
	Events        []EventEntity             `gorm:"foreignKey:OrganizerUserID" json:"events,omitempty"`
	Registrations []EventRegistrationEntity `gorm:"foreignKey:UserID"          json:"registrations,omitempty"`
}

func (UserEntity) TableName() string { return "users" }
