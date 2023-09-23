package middlewares

import (
	"errors"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/pkg/config"
)

func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: handleError,
		SigningKey:   jwtware.SigningKey{Key: []byte(config.TOKEN_KEY)},
		TokenLookup:  "cookie:session_token",
	})
}

func handleError(ctx *fiber.Ctx, err error) error {
	if errors.Is(err, jwtware.ErrJWTMissingOrMalformed) {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "JWT missing/malformed",
		})
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "fail",
		"message": "JWT invalid/expired",
	})
}
