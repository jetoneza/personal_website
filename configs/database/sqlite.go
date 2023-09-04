package database

import "github.com/gofiber/storage/sqlite3/v2"

func Connect(fileName string, tableName string) *sqlite3.Storage {
	store := sqlite3.New(sqlite3.Config{
		Database: fileName,
		Table:    tableName,
	})

	return store
}
