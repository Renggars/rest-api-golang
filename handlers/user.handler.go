package handlers

import (
	"github/database"
	"github/models/entity"
	"github/models/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(c *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	// err := database.DB.Find(&users).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return c.JSON(users)
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	result := database.DB.Debug().Create(&newUser)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": result.Error.Error(),
			"data":    nil},
		)
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}
