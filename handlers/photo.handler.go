package handlers

import (
	"github/database"
	"github/models/entity"
	"github/models/request"
	"github/utils"
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
	filenames := c.Locals("filenames")
	if filenames == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "image cover is required",
		})
	} else {
		filenameData := filenames.([]string)
		for _, filename := range filenameData {
			// Create a new photo entity
			newPhoto := entity.Photo{
				Image:      filename,
				CategoryID: photo.CategoryID,
			}

			errCreatePhoto := database.DB.Debug().Create(&newPhoto).Error
			if errCreatePhoto != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "failed to create photo",
					"error":   errCreatePhoto.Error(),
				})
			}
		}
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func PhotoHandlerDelete(c *fiber.Ctx) error {
	photoId := c.Params("id")

	var photo entity.Photo

	// check availability photo
	err := database.DB.Debug().First(&photo, "id=?", photoId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "photo not found",
		})
	}

	// handle delete photo
	errDeleteFile := utils.HandleRemoveFile(photo.Image)
	if errDeleteFile != nil {
		log.Println("Error Remove File = ", errDeleteFile)
	}

	errDelete := database.DB.Debug().Delete(&photo).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to delete photo",
			"error":   errDelete.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success to delete photo",
	})
}
