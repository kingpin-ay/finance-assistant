package main

import (
	"backend/controller"
	"backend/initializer"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.InitDatabaseConnection()
	initializer.SyncDatabase()
}

func main() {
	server := gin.Default()

	server.POST("/signUp", controller.SignUpUser)
	server.POST("/login", controller.Login)
	server.GET("/validate", middleware.AuthGuard, controller.Validate)

	server.Run()
}
