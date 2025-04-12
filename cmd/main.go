package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"server-go/internal/config"
	"server-go/internal/handlers"
	"server-go/internal/repositories"
	"server-go/internal/routes"
	"server-go/internal/services"

	"github.com/gin-gonic/gin"
)

func InitLogger() {
	logFile, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл для логов: %v", err)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	InitLogger()
	log.Println("Приложение запущено")
	LoadEnv()

	log.Println("Попытка подключения к БД...")
	config.InitDB()

	tourRepository := repositories.NewTourRepository(config.DB)
	tourService := services.NewTourService(tourRepository)
	tourHandler := handlers.NewTourHandler(tourService)

	basketRepository := repositories.NewBasketRepository(config.DB)
	basketService := services.NewBasketService(basketRepository)
	basketHandler := handlers.NewBasketHandler(basketService)

	authRepository := repositories.NewAuthRepository(config.DB)
	authService := services.NewAuthService(authRepository, basketService)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	routes.SetTourRoutes(router, tourHandler)
	routes.SetAuthRoutes(router, authHandler)
	routes.SetBasketRoutes(router, basketHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
