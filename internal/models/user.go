package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string         `json:"username" binding:"required" gorm:"unique"`
	Email     string         `json:"email" binding:"required,email" gorm:"unique"`
	Password  string         `json:"-" binding:"required"`
	Role      string         `json:"role" binding:"required" gorm:"default:'user'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
