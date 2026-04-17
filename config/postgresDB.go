package config

import (
	"commerce-manager/domain/entities"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	var err error
	confString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	DB, err = gorm.Open(postgres.Open(confString), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")

	}
	DB.AutoMigrate(&entities.Transaction{})
	DB.AutoMigrate(&entities.Merchant{})
	DB.AutoMigrate(&entities.User{})
}
