package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go.com/auth/handler" // Ensure the module path is correct
)

// InitializeRoutes sets up authentication and health check routes
func InitializeRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", handler.Register)
		authRoutes.POST("/login", handler.Login)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Healthy"})
	})
}

