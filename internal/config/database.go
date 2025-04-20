package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
		return fmt.Errorf("недостаточно данных для подключения к базе данных: проверьте переменные окружения DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME")
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslMode)

	sqlDB, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("ошибка при открытии соединения: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("БД недоступна: %v", err)
	}

	if err := goose.Up(sqlDB, "./internal/config/migrations"); err != nil {
		return fmt.Errorf("ошибка при применении миграций: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("ошибка при инициализации GORM: %v", err)
	}

	DB = gormDB

	log.Println("База данных инициализирована и миграции успешно применены")
	return nil
}
