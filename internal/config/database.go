package config

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "travel_agency"
	dbUser := "postgres"
	dbPass := "Kz123456"
	dbPort := "5432"
	sslmode := "disable"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	if err := goose.Up(sqlDB, "./internal/config/migrations"); err != nil {
		log.Fatalf("Ошибка при применении миграций: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при инициализации GORM: ", err)
	}

	DB = gormDB
	log.Println("База данных инициализирована и миграции успешно применены")
}
