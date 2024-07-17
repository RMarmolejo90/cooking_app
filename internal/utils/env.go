package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("\n---\tError loading environment variables!!!\t---\n")
	}
}
