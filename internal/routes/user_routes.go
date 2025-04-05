package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
)

func SetAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}
