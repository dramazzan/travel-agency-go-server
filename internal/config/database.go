package config

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

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

	// Удаляем defer sqlDB.Close() — потому что GORM будет использовать это соединение!

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("БД недоступна: %v", err)
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
