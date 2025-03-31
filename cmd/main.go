package main

import (
	"log"
	"server-go/internal/config"
	"server-go/internal/handlers"
	"server-go/internal/models"
	"server-go/internal/repositories"
	"server-go/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	err = db.AutoMigrate(&models.Tour{})
	if err != nil {
		log.Fatalf("Не удалось выполнить миграцию: %v", err)
	}

	tourRepository := repositories.NewTourRepository(db)
	tourService := services.NewTourService(tourRepository)
	tourHandler := handlers.NewTourHandler(tourService)

	router := gin.Default()

	api := router.Group("/api")
	{
		tours := api.Group("/tours")
		{
			tours.GET("", tourHandler.GetAllTours)
			tours.GET("/:id", tourHandler.GetTourByID)
			tours.POST("", tourHandler.CreateTour)
			tours.PUT("/:id", tourHandler.UpdateTour)
			tours.DELETE("/:id", tourHandler.DeleteTour)
		}
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
