package database

import (
	"log"

	"api-rest-training/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=api_rest_go port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Impossible de connecter à la base de donnée")
	}

	DB = db

	DB.AutoMigrate(&models.User{})
}
