package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
)

func SetTourRoutes(router *gin.Engine, tourHandler *handlers.TourHandler) {
	tours := router.Group("/tours")
	{
		tours.GET("", tourHandler.GetAllTours)
		tours.GET("/:id", tourHandler.GetTourByID)
		tours.POST("", tourHandler.CreateTour)
		tours.PUT("/:id", tourHandler.UpdateTour)
		tours.DELETE("/:id", tourHandler.DeleteTour)
	}

}
