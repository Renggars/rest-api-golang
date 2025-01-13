package handlers

import (
	"fmt"
	"github/database"
	"github/models/entity"
	"github/models/request"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(c *fiber.Ctx) error {
	book := new(request.BookCreateRequest)
	if err := c.BodyParser(book); err != nil {
		return err
	}

	// validate
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Handle File
	file, errFile := c.FormFile("cover")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	var filename string
	if file != nil {
		filename = file.Filename

		errSaveFile := c.SaveFile(file, fmt.Sprint("public/covers/", file.Filename))
		if errSaveFile != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "failed to save file",
				"error":   errSaveFile.Error(),
			})
		}
	} else {
		log.Println("no file uploaded")
	}

	// Create a new book entity
	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filename,
	}

	errCreateBook := database.DB.Debug().Create(&newBook).Error
	if errCreateBook != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to create book",
			"error":   errCreateBook.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})
}
