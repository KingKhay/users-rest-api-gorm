package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"users-rest-api-gorm/controller"
	"users-rest-api-gorm/initializers"
	"users-rest-api-gorm/middleware"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {

	router := gin.Default()

	group := router.Group("/users")

	group.GET("", controller.GetAllUsers)

	group.GET("/:id", controller.GetUserById)

	group.POST("", controller.CreateUser)

	group.PUT("/:id", controller.UpdateUser)

	group.DELETE("/:id", controller.DeleteUser)

	router.POST("/login", controller.Login)

	router.GET("/hello", middleware.JwtAuthFilter, func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})

	router.Run(":9400")
}
