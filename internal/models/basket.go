package models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint   `json:"userid" gorm:"unique"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Tours  []Tour `json:"tours" gorm:"many2many:basket_tours;"`
}
