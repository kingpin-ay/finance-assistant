package main

import (
	"backend/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.InitDatabaseConnection()
}

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run()
}
