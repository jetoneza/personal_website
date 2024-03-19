package database

import (
	"errors"
	"log"
	"time"

	"github.com/jetoneza/personal_website/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(name string) (*gorm.DB, error) {
	// This is for sqlite only
	dbPath := "db/"

	db, err := gorm.Open(sqlite.Open(dbPath+name), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, errors.New("Failed opening connection to sqlite:" + err.Error())
	}

	log.Println("Running migrations.")
	db.AutoMigrate(&models.Post{}, &models.User{}, &models.Event{}, &models.Integration{}, &models.UserSetting{})

	log.Println("Database successfully connected.")

	return db, err
}
