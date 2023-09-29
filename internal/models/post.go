package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jetoneza/personal_website/pkg/utils"
	"gorm.io/gorm"
)

type Post struct {
	ID              string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Title           string    `gorm:"type:varchar(255);uniqueIndex:idx_post_title,LENGTH(255);not null" json:"title,omitempty"`
	Slug            string    `gorm:"type:varchar(255);uniqueIndex:idx_post_slug,LENGTH(255);not null" json:"slug,omitempty"`
	Description     string    `gorm:"type:text" json:"description,omitempty"`
	Content         string    `gorm:"type:text;not null" json:"content,omitempty"`
	Category        string    `gorm:"type:varchar(100)" json:"category,omitempty"`
	MetaTitle       string    `gorm:"type:varchar(255)" json:"meta_title,omitempty"`
	MetaDescription string    `gorm:"type:text" json:"meta_description,omitempty"`
	MetaKeywords    string    `gorm:"type:varchar(255)" json:"meta_keywords,omitempty"`
	Published       bool      `gorm:"default:false;not null" json:"published"`
	CreatedAt       time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt       time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	post.ID = uuid.New().String()
	post.CreatedAt = now
	post.UpdatedAt = now

	return nil
}

func (post *Post) ConvertContentToHtml() {
	md := []byte(post.Content)
	htmlBytes := utils.MdToHtml(md)
	post.Content = string(htmlBytes)
}
