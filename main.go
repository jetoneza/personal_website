package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jetoneza/personal_website/pkg/application"
	"github.com/jetoneza/personal_website/web"
)

const (
	appName = "Jet Ordaneza Personal Website"
	port    = ":3000"
)

func main() {
	app := application.New()
	app.ConnectDB()

	fiberApp := fiber.New(fiber.Config{
		AppName: appName,
	})
	defer fiberApp.Shutdown()

	// Middlewares
	fiberApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// TODO: move to routes
	fiberApp.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "jetrooper.me API",
		})
	})

	// Serve static files
	fiberApp.All("/*", filesystem.New(filesystem.Config{
		Root:         web.Build(),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))

	// Go!
	if err := fiberApp.Listen(port); err != nil {
		panic(err)
	}
}
