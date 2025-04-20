package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
	"server-go/internal/middleware"
)

func SetTourRoutes(r *gin.Engine, h *handlers.TourHandler) {
	tours := r.Group("/tours")
	{
		tours.GET("", h.GetAllTours)
		tours.GET("/:id", h.GetTourByID)
	}

	admin := r.Group("/admin/tour", middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("", h.CreateTour)
		admin.PUT("/:id", h.UpdateTour)
		admin.DELETE("/:id", h.DeleteTour)
	}
}
