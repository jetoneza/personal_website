package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Integration struct {
	ID           string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name         string    `gorm:"type:varchar(255)" json:"name,omitempty"`
	TokenType    string    `gorm:"type:varchar(255)" json:"token_type,omitempty"`
	ExpiresAt    int       `gorm:"not null" json:"expires_at"`
	ExpiresIn    int       `gorm:"not null" json:"expires_in"`
	RefreshToken string    `gorm:"not null" json:"refresh_token"`
	AccessToken  string    `gorm:"not null" json:"access_token"`
	CreatedAt    time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

func (integration *Integration) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	integration.ID = uuid.New().String()
	integration.CreatedAt = now
	integration.UpdatedAt = now

	return nil
}
