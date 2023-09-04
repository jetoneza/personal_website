package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jetoneza/personal_website/configs/database"
	"github.com/jetoneza/personal_website/web"
)

const (
	appName = "Jet Ordaneza Personal Website"
	port    = ":3000"
)

func main() {
	// TODO: Handle panics
	db := database.Connect("./posts.sqlite3", "posts")
	defer db.Close()

	app := fiber.New(fiber.Config{
		AppName: appName,
	})
	defer app.Shutdown()

	// Middlewares
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Serve static files
	app.All("/*", filesystem.New(filesystem.Config{
		Root:         web.Build(),
		NotFoundFile: "index.html",
		Index:        "index.html",
	}))

	// Go!
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
