package routes

import (
	"github/config"
	"github/handlers"
	"github/middlewares"
	"github/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	app.Static("/public", config.ProjectRootPath+"/public/assets")

	app.Post("/login", handlers.LoginHandler)

	app.Post("/user", middlewares.AuthMiddleware, handlers.UserHandlerCreate)
	app.Get("/user", middlewares.AuthMiddleware, handlers.UserHandlerGetAll)
	app.Get("/user/:id", middlewares.AuthMiddleware, handlers.UserHandlerGetById)
	app.Put("/user/:id", middlewares.AuthMiddleware, handlers.UserHandlerUpdate)
	app.Put("/user/:id/update-email", middlewares.AuthMiddleware, handlers.UserHandlerUpdateEmail)
	app.Delete("/user/:id", middlewares.AuthMiddleware, handlers.UserHandlerDelete)

	app.Post("/book", utils.HandleSingleFile, handlers.BookHandlerCreate)
}
