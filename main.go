package main

import (
	"fmt"
	"log"

	"api-rest-training/database"
	"api-rest-training/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Fichier .env absent, utilisation des variables système")
	}

	database.Connect()
	router := gin.Default()
	routes.Setup(router)
	router.Run(":8080")
	fmt.Println("Let's go for the Crud !!!")
}
