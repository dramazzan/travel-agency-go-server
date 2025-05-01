package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"server-go/internal/config"
	"server-go/internal/handlers"
	"server-go/internal/repositories"
	"server-go/internal/routes"
	"server-go/internal/services"
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
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("Локально: .env не найден")
		}
	}
}

func main() {
	InitLogger()
	log.Println("Приложение запущено")
	LoadEnv()

	log.Println("Попытка подключения к БД...")
	config.InitDB()

	tourRepository := repositories.NewTourRepository(config.DB)
	basketRepository := repositories.NewBasketRepository(config.DB)
	authRepository := repositories.NewAuthRepository(config.DB)

	tourService := services.NewTourService(tourRepository)
	basketService := services.NewBasketService(basketRepository)
	authService := services.NewAuthService(authRepository, basketService)

	tourHandler := handlers.NewTourHandler(tourService)
	basketHandler := handlers.NewBasketHandler(basketService)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetTourRoutes(router, tourHandler)
	routes.SetAuthRoutes(router, authHandler)
	routes.SetBasketRoutes(router, basketHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
