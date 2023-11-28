package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jetoneza/personal_website/internal/handlers"
	"github.com/jetoneza/personal_website/internal/middlewares"
)

func EventRoutes(router fiber.Router, handlers *handlers.Handler) {
	r := router.Group("/events")

	r.Get("/", middlewares.AuthMiddleware(), handlers.GetAllEvents)
	r.Post("/", middlewares.AuthMiddleware(), handlers.CreateEvent)

	r.Get("/:id", middlewares.AuthMiddleware(), handlers.GetEvent)
	r.Delete("/:id", middlewares.AuthMiddleware(), handlers.DeleteEvent)
	r.Put("/:id", middlewares.AuthMiddleware(), handlers.UpdateEvent)
}
