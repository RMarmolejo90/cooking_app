package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
		return err
	}
	return nil
}
