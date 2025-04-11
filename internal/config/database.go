package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

var DB *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbName := "travel_agency"
	dbUser := "postgres"
	dbPass := "Kz123456"
	dbPort := "5432"
	sslmode := "disable"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	var sqlDB *sql.DB
	var err error

	// Попробовать подключение несколько раз с паузой
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		sqlDB, err = sql.Open("postgres", dbUrl)
		if err == nil && sqlDB.Ping() == nil {
			break
		}
		log.Printf("Попытка %d: БД еще не готова. Жду 2 секунды...\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
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
