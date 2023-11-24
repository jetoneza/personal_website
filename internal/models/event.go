package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Title     string    `gorm:"type:varchar(255);" json:"title,omitempty"`
	AllDay    bool      `gorm:"default:false;not null" json:"all_day"`
	Start     time.Time `json:"start,omitempty"`
	End       time.Time `json:"end,omitempty"`
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

func (event *Event) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	event.ID = uuid.New().String()
	event.CreatedAt = now
	event.UpdatedAt = now

	return nil
}
