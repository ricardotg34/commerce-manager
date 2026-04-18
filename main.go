package main

import (
	"commerce-manager/config"
	"commerce-manager/presentation"
	"commerce-manager/presentation/middlewares"
	"database/sql"
	"io"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		println("Failed to load env variables")
	}

	config.ConnectDatabase()

	server := gin.New()
	server.Use(gin.Recovery())

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(file, os.Stdout), nil))

	server.Use(middlewares.LoggerMiddleware(logger))
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	presentation.RegisterRoutes(server)

	server.Run() // listens on 0.0.0.0:8080 by default
}
