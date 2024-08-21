package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	//read env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
