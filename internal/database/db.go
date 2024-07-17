package database

import (
	"fmt"
	"log"
	"os"

	"github.com/RMarmolejo90/cooking_app/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	utils.LoadEnv()
	user, password, dbname := os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DATABASE")
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=9920 sslmode=disable", user, password, dbname)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Printf("\n\n---\tConnected To The Database\t---\n\n")

	// Need to add all models for migration
	err = DB.AutoMigrate()
	if err != nil {
		log.Print(err.Error())
	}

}
