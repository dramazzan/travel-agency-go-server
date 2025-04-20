package routes

import (
	"github.com/gin-gonic/gin"
	"server-go/internal/handlers"
	"server-go/internal/middleware"
)

func SetAuthRoutes(r *gin.Engine, h *handlers.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	user := r.Group("/user", middleware.AuthMiddleware())
	{
		user.GET("/dashboard", h.GetUserData)
		user.PUT("/update", h.UpdateUserData)
	}

	// admin := r.Group("/admin", middleware.AuthMiddleware(), middleware.AdminMiddleware())
	// {
	// 	admin.GET("/dashboard", h.OpenAdminProfile)
	// }

	// protected := r.Group("/protected", middleware.AuthMiddleware())
	// {
	// 	protected.GET("/dashboard", h.OpenUserProfile)
	// }
}
