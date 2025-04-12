package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
	"server-go/internal/middleware"
)

func SetBasketRoutes(router *gin.Engine, basketHandler *handlers.BasketHandler) {
	basketRoute := router.Group("/basket")
	basketRoute.Use(middleware.AuthMiddleware())
	{
		basketRoute.POST("/tours/:tourID", basketHandler.AddTourOnBasket)
		basketRoute.GET("/tours", basketHandler.GetBasket)
		basketRoute.DELETE("/tours/:tourID", basketHandler.RemoveTourFromBasket)
	}
}
