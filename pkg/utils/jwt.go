package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jetoneza/personal_website/pkg/config"
)

type AuthTokenPayload struct {
	ID    string
	Email string
}

func GenerateAuthToken(payload *AuthTokenPayload) string {
	duration, err := time.ParseDuration(config.TOKEN_EXPIRY)
	if err != nil {
		panic("Invalid time duration.")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   time.Now().Add(duration).Unix(),
		"id":    payload.ID,
		"email": payload.Email,
	})

	token, err := t.SignedString([]byte(config.TOKEN_KEY))
	if err != nil {
		panic(err)
	}

	return token
}
