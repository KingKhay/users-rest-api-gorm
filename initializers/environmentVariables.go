package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvironmentVariables() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
