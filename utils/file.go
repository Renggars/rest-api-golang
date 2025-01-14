package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/covers/"

func HandleSingleFile(c *fiber.Ctx) error {
	// Handle File
	file, errFile := c.FormFile("cover")
	if errFile != nil {
		log.Println("Error File = ", errFile)
	}

	var filename *string
	if file != nil {
		errChekoContentType := CheckContentType(file, "image/jpeg", "image/png", "image/jpg")
		if errChekoContentType != nil {
			return c.Status(422).JSON(fiber.Map{
				"message": errChekoContentType.Error(),
			})
		}

		filename = &file.Filename
		extensionFile := filepath.Ext(*filename)
		newFilename := fmt.Sprintf("gambar-satu%s", extensionFile)

		errSaveFile := c.SaveFile(file, fmt.Sprint("public/covers/", newFilename))
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
			extensionFile := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%d-%s%s", i, "gambar", extensionFile)

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

func HandleRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("Error Remove File = ", err)
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Error Remove File = ", err)
			return err
		}
	}

	return nil
}

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}

		return errors.New("file must be an image")
	} else {
		return errors.New("not found content type to be checked")
	}
}
