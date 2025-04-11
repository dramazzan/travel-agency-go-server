package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

var DB *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := "disable"

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Ошибка при открытии соединения: %v", err)
	}
	defer sqlDB.Close()

	// Проверка соединения
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("БД недоступна: %v", err)
	}

	// Применение миграций
	if err := goose.Up(sqlDB, "./internal/config/migrations"); err != nil {
		log.Fatalf("Ошибка при применении миграций: %v", err)
	}

	// Инициализация GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка при инициализации GORM: ", err)
	}

	DB = gormDB
	log.Println("База данных инициализирована и миграции успешно применены")
}
