package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type User struct {
	gorm.Model

	Name  string
	Email string
	Age   int
}

func main() {
	host := os.Getenv("host")
	dbUser := os.Getenv("dbUser")
	dbName := os.Getenv("dbName")
	dbPassword := os.Getenv("dbPassword")
	dbPort := os.Getenv("dbPort")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable  port=%s", host, dbUser, dbPassword, dbName, dbPort)

	//Open Connection to database with Gorm
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

}
