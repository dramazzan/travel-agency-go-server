package models

import (
	"time"

	"gorm.io/gorm"
)

type Tour struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	Destination string         `json:"destination" binding:"required"`
	StartDate   time.Time      `json:"start_date" binding:"required"`
	EndDate     time.Time      `json:"end_date" binding:"required"`
	Price       float64        `json:"price" binding:"required" gorm:"check:price >= 0"`
	MaxCapacity int            `json:"max_capacity"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
