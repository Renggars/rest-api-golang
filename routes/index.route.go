package routes

import (
	"github/config"
	"github/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Static("/public", config.ProjectRootPath+"/public/assets")
	app.Post("/user", handlers.UserHandlerCreate)
	app.Get("/user", handlers.UserHandlerGetAll)
	app.Get("/user/:id", handlers.UserHandlerGetById)
	app.Put("/user/:id", handlers.UserHandlerUpdate)
	app.Put("/user/:id/update-email", handlers.UserHandlerUpdateEmail)
	app.Delete("/user/:id", handlers.UserHandlerDelete)
}
