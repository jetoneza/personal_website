package application

import (
	"log"

	"github.com/jetoneza/personal_website/pkg/database"
	"gorm.io/gorm"
)

type Application struct {
	db *gorm.DB
}

func New() *Application {
	return &Application{}
}

func (a *Application) ConnectDB() *gorm.DB {
	if a.db != nil {
		return a.db
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	a.db = db

	return db
}
