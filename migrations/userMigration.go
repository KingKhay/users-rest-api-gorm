package main

import (
	"log"
	"users-rest-api-gorm/initializers"
	"users-rest-api-gorm/models"
)

func init() {
	initializers.ConnectToDatabase()
}

func main() {

	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
