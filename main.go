package main

import (
	"github.com/gin-gonic/gin"
	"my-gin-jwt-server/handlers"
	"my-gin-jwt-server/middleware"
)

func main() {
	router := gin.Default()

	// Public routes
	router.POST("/login", handlers.LoginHandler)

	// Protected routes
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("", handlers.ProtectedHandler)

	router.Run(":8080")
}
