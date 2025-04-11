package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
	"server-go/internal/middleware"
)

func SetAuthRoutes(router *gin.Engine, authHandler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(200, gin.H{"message": "Welcome Admin " + username + "!"})
		})
	}

	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/dashboard", func(c *gin.Context) {
			username := c.GetString("username")
			c.JSON(200, gin.H{
				"message": "Welcome to your dashboard, " + username,
			})
		})
	}

}
