package main

import (
	"log"
	"server-go/internal/config"
	"server-go/internal/handlers"
	"server-go/internal/repositories"
	"server-go/internal/routes"
	"server-go/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	config.InitDB()

	tourRepository := repositories.NewTourRepository(config.DB)
	tourService := services.NewTourService(tourRepository)
	tourHandler := handlers.NewTourHandler(tourService)

	router := gin.Default()
	routes.SetupRoutes(router, tourHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
