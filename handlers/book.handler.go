package handlers

import (
	"fmt"
	"github/database"
	"github/models/entity"
	"github/models/request"

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

	// Validation Required Image Cover
	var filenameString string
	filename := c.Locals("filename")
	if filename == nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "image cover is required",
		})
	} else {
		filenameString = fmt.Sprintf("(%v)", filename)
	}

	// Create a new book entity
	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
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
