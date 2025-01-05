package routes

import (
	"github/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/user", handlers.UserHandlerGetAll)
	app.Post("/user", handlers.UserHandlerCreate)
}
