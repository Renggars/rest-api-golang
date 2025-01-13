package handlers

import (
	"fmt"
	"github/models/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PhotoHandlerCreate(c *fiber.Ctx) error {
	photo := new(request.PhotoCreateRequest)
	if err := c.BodyParser(photo); err != nil {
		return err
	}

	// validate
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Validation Required Image Cover
	var filenameString string

	filenames := c.Locals("filenames")
	if filenames == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "image cover is required",
		})
	} else {
		filenameString = fmt.Sprintf("(%v)", filenames)
	}

	log.Println(filenameString)

	// Create a new photo entity
	// newPhoto := entity.Book{
	// 	Image:      filename,
	// 	CategoryID: photo.CategoryID,
	// }

	// errCreatePhoto := database.DB.Debug().Create(&newPhoto).Error
	// if errCreatePhoto != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"message": "failed to create photo",
	// 		"error":   errCreatePhoto.Error(),
	// 	})
	// }

	return c.JSON(fiber.Map{
		"message": "success",
		// "data":    newPhoto,
	})
}
