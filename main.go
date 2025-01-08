package main

import (
	"github/database"
	"github/database/migration"

	// "github/database/migration"
	"github/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	database.DatabaseInit()

	// Run migration
	migration.RunMigration()

	app := fiber.New()

	// Initialize routes
	routes.RouteInit(app)

	app.Listen(":3000")
}
