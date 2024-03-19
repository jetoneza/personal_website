package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          string      `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name        string      `gorm:"not null"`
	Email       string      `gorm:"uniqueIndex;not null"`
	Password    string      `gorm:"not null"`
	UserSetting UserSetting `gorm:"foreignKey:UserID;references:ID" json:"user_setting,omitempty"`
	CreatedAt   time.Time   `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt   time.Time   `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	user.ID = uuid.New().String()
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}
