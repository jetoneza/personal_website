package config

import (
	"fmt"
	"os"
  _ "github.com/joho/godotenv/autoload"
)

var (
	APP_NAME = getEnv("APP_NAME", "Jet Ordaneza Personal Website")
	DB       = getEnv("DB", "jetdb.sqlite")
	PORT     = getEnv("PORT", "3000")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
    fmt.Println(value)
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
