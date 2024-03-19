package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSetting struct {
	ID          string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	UserID      string    `gorm:"type:char(36);not null" json:"user_id,omitempty"`
	Name        string    `gorm:"type:char(36)" json:"name,omitempty"`
	Value       string    `gorm:"type:text" json:"name,omitempty"`
	Description string    `gorm:"type:varchar(255)" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

func (setting *UserSetting) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	setting.ID = uuid.New().String()
	setting.CreatedAt = now
	setting.UpdatedAt = now

	return nil
}
