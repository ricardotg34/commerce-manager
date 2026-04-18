package main

import (
	"commerce-manager/config"
	_ "commerce-manager/docs"
	"commerce-manager/presentation"
	"commerce-manager/presentation/middlewares"
	"database/sql"
	"io"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title			Swagger Example API
// @version		1.0
// @description	API that manages merchants and its transactions.
// @host		localhost:8080
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
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	presentation.RegisterRoutes(server)

	server.Run()
}
