package migration

import (
	"fmt"
	"github/database"
	"github/models/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
