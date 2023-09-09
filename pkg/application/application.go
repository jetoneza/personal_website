package application

import (
	"log"

	"github.com/jetoneza/personal_website/pkg/database"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func New() *Application {
	return &Application{}
}

func (a *Application) InitializeDB(name string) *gorm.DB {
	if a.DB != nil {
		return a.DB
	}

	db, err := database.Connect(name)
	if err != nil {
		log.Fatal(err)
	}

	a.DB = db

	return db
}
