package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"-" gorm:"column:password"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
