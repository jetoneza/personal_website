package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(name string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Failed opening connection to sqlite:" + err.Error())
	}

	// TODO: Add auto migrate

	return db, err
}
