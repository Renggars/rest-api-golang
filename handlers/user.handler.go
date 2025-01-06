package handlers

import (
	"github/database"
	"github/models/entity"
	"github/models/request"
	"github/models/responses"
	"log"

	"github.com/go-playground/validator/v10"
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

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil},
		)
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

func UserHandlerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User
	result := database.DB.Debug().First(&user, userId)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": result.Error.Error(),
			"data":    nil},
		)
	}

	userResponse := responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}

func UserHandlerUpdate(c *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user entity.User

	userId := c.Params("id")

	result := database.DB.Debug().First(&user, userId)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	if userRequest.Address != "" {
		user.Address = userRequest.Address
	}

	if userRequest.Phone != "" {
		user.Phone = userRequest.Phone
	}

	resultUpdate := database.DB.Debug().Save(&user)
	if resultUpdate.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": resultUpdate.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdateEmail(c *fiber.Ctx) error {
	userRequest := new(request.UserUpdateEmailRequest)
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Membuat instance validator baru
	validate := validator.New()

	// Melakukan validasi terhadap body request
	if err := validate.Struct(userRequest); err != nil {
		// Mengembalikan error validasi
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var user entity.User

	userId := c.Params("id")

	result := database.DB.Debug().First(&user, userId)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//check availability email
	checkEmail := database.DB.Debug().First(&user, "email = ?", userRequest.Email)
	if checkEmail.Error == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "email already exist",
		})
	}

	user.Email = userRequest.Email

	resultUpdate := database.DB.Debug().Save(&user)
	if resultUpdate.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": resultUpdate.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User
	//check availability user
	result := database.DB.Debug().First(&user, userId)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	resultDelete := database.DB.Debug().Delete(&user)
	if resultDelete.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": resultDelete.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
