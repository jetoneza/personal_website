package main

import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/fiber/v2/middleware/logger"

const (
	appName = "Jet Ordaneza Personal Website"
	port    = ":3000"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: appName,
	})
	defer app.Shutdown()

  // Middlewares
  app.Use(logger.New(logger.Config{
    Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
  }))

	// Go!
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
