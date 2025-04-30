package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
	"server-go/internal/middleware"
)

func SetTourRoutes(router *gin.Engine, tourHandler *handlers.TourHandler) {
	tours := router.Group("/tours")
	{
		tours.GET("", tourHandler.GetAllTours)
		tours.GET("/:id", tourHandler.GetTourByID)
		tours.GET("/category", tourHandler.GetTourByCategory)
	}

	adminRoutes := router.Group("/admin/tour")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{

		adminRoutes.POST("", tourHandler.CreateTour)
		adminRoutes.PUT("/:id", tourHandler.UpdateTour)
		adminRoutes.DELETE("/:id", tourHandler.DeleteTour)
	}

}
