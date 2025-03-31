package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Kz123456 dbname=travel_agency port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
