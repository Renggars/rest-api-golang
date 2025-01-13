package entity

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	Image      string         `json:"image"`
	CategoryID int            `json:"category_id" gorm:"column:category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
