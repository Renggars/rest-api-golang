package main

import (
	"github/database"
	"github/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	database.DatabaseInit()

	app := fiber.New()

	// Initialize routes
	routes.RouteInit(app)

	app.Listen(":3000")
}
