package main

import (
	"backend/controller"
	"backend/initializer"
	"backend/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.InitDatabaseConnection()
	initializer.SyncDatabase()
}

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Allow requests from this origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Content-Type", "Authorization"},           // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Headers exposed to the client
		AllowCredentials: true,                                                // Allow cookies or other credentials
		MaxAge:           12 * time.Hour,                                      // Cache preflight responses
	}))

	server.POST("/auth/signUp", controller.SignUpUser)
	server.POST("/auth/login", controller.Login)
	server.GET("/auth/validate", middleware.AuthGuard, controller.Validate)

	server.Run()
}
