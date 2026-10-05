package routes

import (
	"api-rest-training/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
