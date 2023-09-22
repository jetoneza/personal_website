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

type AuthClaims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAuthToken(payload *AuthTokenPayload) (string, time.Time) {
	duration, err := time.ParseDuration(config.TOKEN_EXPIRY)
	if err != nil {
		panic("Invalid time duration.")
	}

	expiry := time.Now().Add(duration)

	claims := AuthClaims{
		payload.ID,
		payload.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.APP_NAME,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(config.TOKEN_KEY))
	if err != nil {
		panic(err)
	}

	return token, expiry
}
