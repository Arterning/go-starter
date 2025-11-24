package handlers

import (
	"go-starter/internal/middleware"
	"go-starter/pkg/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *UserHandler, jwtManager *utils.JWTManager) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/register", userHandler.Register)
				auth.POST("/login", userHandler.Login)
			}

			users := v1.Group("/users")
			users.Use(middleware.AuthMiddleware(jwtManager))
			{
				users.GET("/profile", userHandler.GetProfile)
			}
		}
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	return router
}
