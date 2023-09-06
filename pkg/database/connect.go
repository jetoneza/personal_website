package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

  // TODO: Add auto migrate

	return db, err
}
