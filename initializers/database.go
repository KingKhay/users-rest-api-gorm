package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func ConnectToDatabase() {
	host := os.Getenv("host")
	dbUser := os.Getenv("dbuser")
	dbName := os.Getenv("dbname")
	dbPassword := os.Getenv("dbpassword")
	dbPort := os.Getenv("dbport")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable  port=%s", host, dbUser, dbPassword, dbName, dbPort)

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
