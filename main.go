package main

import (
	"fmt"

	"api-rest-training/database"
	"api-rest-training/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	router := gin.Default()
	routes.Setup(router)
	router.Run(":8080")
	fmt.Println("Let's go for the Crud !!!")
}
