package main

import "github.com/gofiber/fiber/v2"

const (
	appName = "Jet Ordaneza Personal Website"
	port    = ":3000"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: appName,
	})
	defer app.Shutdown()

	// Go!
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
