package handlers

import (
	"github/database"
	"github/models/entity"
	"github/models/request"
	"github/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": errValidate.Error(),
		})
	}

	// check user availability
	var user entity.User
	result := database.DB.Debug().First(&user, "email = ?", loginRequest.Email).Error
	if result != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// check password
	isValid := utils.ComparePassword(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid password",
		})
	}

	return c.JSON(fiber.Map{
		"token": "secret-token",
	})
}
