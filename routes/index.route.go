package routes

import (
	"github/config"
	"github/handlers"

	"github.com/gofiber/fiber/v2"
)

func middleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token != "secret-token" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	return c.Next()
}

func RouteInit(app *fiber.App) {
	app.Static("/public", config.ProjectRootPath+"/public/assets")
	app.Post("/user", middleware, handlers.UserHandlerCreate)
	app.Get("/user", middleware, handlers.UserHandlerGetAll)
	app.Get("/user/:id", middleware, handlers.UserHandlerGetById)
	app.Put("/user/:id", middleware, handlers.UserHandlerUpdate)
	app.Put("/user/:id/update-email", middleware, handlers.UserHandlerUpdateEmail)
	app.Delete("/user/:id", middleware, handlers.UserHandlerDelete)
}
