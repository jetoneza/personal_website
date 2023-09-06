package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Title     string    `gorm:"type:varchar(255);uniqueIndex:idx_notes_title,LENGTH(255);not null" json:"title,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	Category  string    `gorm:"varchar(100)" json:"category,omitempty"`
	Published bool      `gorm:"default:false;not null" json:"published"`
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.ID = uuid.New().String()

	return nil
}
