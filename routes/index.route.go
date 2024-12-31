package routes

import (
	"github/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Get("/", controllers.UserHandlerRead)
}
