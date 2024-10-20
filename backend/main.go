package main

import (
	"backend/controller"
	"backend/initializer"

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

	server.Run()
}
