package config

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	APP_NAME      = getEnv("APP_NAME", "Jet Ordaneza Personal Website")
	DB            = getEnv("DB", "db.sqlite3")
	PORT          = getEnv("PORT", "3000")
	SVELTE_PORT   = getEnv("SVELTE_PORT", "3003")
	TOKEN_KEY     = getEnv("TOKEN_KEY", "")
	TOKEN_EXPIRY  = getEnv("TOKEN_EXPIRY", "10h")
	S3_BUCKET     = getEnv("S3_BUCKET", "")
	S3_OBJECT_KEY = getEnv("S3_OBJECT_KEY", "")
	AWS_REGION    = getEnv("AWS_REGION", "")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
