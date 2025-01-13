package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(c *fiber.Ctx) error {
	// Handle File
	file, errFile := c.FormFile("cover")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	var filename *string
	if file != nil {
		filename = &file.Filename

		errSaveFile := c.SaveFile(file, fmt.Sprint("public/covers/", *filename))
		if errSaveFile != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "failed to save file",
				"error":   errSaveFile.Error(),
			})
		}
	} else {
		log.Println("no file uploaded")
	}

	if filename != nil {
		c.Locals("filename", *filename)
	} else {
		c.Locals("filename", nil)
	}

	return c.Next()
}

func HandleMultipleFIle(c *fiber.Ctx) error {
	form, errForm := c.MultipartForm()
	if errForm != nil {
		log.Println("Error Form = ", errForm)
	}

	files := form.File["photos"]

	var filenames []string

	for i, file := range files {
		var filename string
		if file != nil {
			filename = fmt.Sprintf("%d-%v", i, file.Filename)

			errSaveFile := c.SaveFile(file, fmt.Sprint("public/covers/", filename))
			if errSaveFile != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "failed to save file",
					"error":   errSaveFile.Error(),
				})
			}
		} else {
			log.Println("no file uploaded")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	c.Locals("filenames", filenames)

	return c.Next()
}
