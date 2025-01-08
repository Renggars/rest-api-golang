package middlewares

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("x-token")
	if token != "secret-token" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {
	return c.Next()
}
