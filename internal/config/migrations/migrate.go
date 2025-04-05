package migrations

import (
	"gorm.io/gorm"
	"log"
	"server-go/internal/models"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(models.Models()...)
	if err != nil {
		log.Fatalf("\033[31mНе удалось выполнить миграцию: %v\033[0m", err)
	}

	log.Println("\033[32mМиграция прошла успешно\033[0m")
}
