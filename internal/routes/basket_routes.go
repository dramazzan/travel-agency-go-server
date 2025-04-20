package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
	"server-go/internal/middleware"
)

func SetBasketRoutes(r *gin.Engine, h *handlers.BasketHandler) {
	basket := r.Group("/basket", middleware.AuthMiddleware())
	{
		basket.POST("/tours/:tourID", h.AddTourToBasket)
		basket.GET("/tours", h.GetBasket)
		basket.DELETE("/tours/:tourID", h.RemoveTourFromBasket)
	}
}
