package main

import (
	"commerce-manager/config"
	"commerce-manager/presentation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		println("Failed to load env variables")
	}

	config.ConnectDatabase()

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	presentation.RegisterRoutes(server)

	server.Run() // listens on 0.0.0.0:8080 by default
}
