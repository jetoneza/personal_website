package database

import (
	"errors"
	"log"

	"github.com/jetoneza/personal_website/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(name string) (*gorm.DB, error) {
  // This is for sqlite only
	dbPath := "db/"

	db, err := gorm.Open(sqlite.Open(dbPath+name), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Failed opening connection to sqlite:" + err.Error())
	}

	log.Println("Running migrations.")
	db.AutoMigrate(&models.Post{})

	log.Println("Database successfully connected.")

	return db, err
}
