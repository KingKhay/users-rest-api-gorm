package main

import (
	"github.com/gin-gonic/gin"
	"users-rest-api-gorm/controller"
	"users-rest-api-gorm/initializers"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {

	router := gin.Default()

	router.GET("/users", controller.GetAllUsers)

	router.GET("/users/:id", controller.GetUserById)

	router.POST("/users", controller.CreateUser)

	router.PUT("/users/:id", controller.UpdateUser)

	router.DELETE("/users/:id", controller.DeleteUser)

	router.Run(":9400")
}
