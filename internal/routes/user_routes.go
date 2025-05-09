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

	dashboard := router.Group("/user")
	dashboard.Use(middleware.AuthMiddleware())
	{
		dashboard.GET("/dashboard", authHandler.GetUserData)
		dashboard.PUT("/update", authHandler.UpdateUserData)

	}

	admin := router.Group("/admin/user")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/all", authHandler.GetAllUsers)
		admin.PUT("/block/:id", authHandler.BlockUser)
		admin.PUT("/unblock/:id", authHandler.UnblockUser)
		admin.DELETE("/:id", authHandler.DeleteUser)

	}

	//admin := router.Group("/admin")
	//admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	//{
	//	admin.GET("/dashboard", authHandler.OpenAdminProfile)
	//}

	//protected := router.Group("/protected")
	//protected.Use(middleware.AuthMiddleware())
	//{
	//	protected.GET("/dashboard", authHandler.OpenUserProfile)
	//
	//}

}
