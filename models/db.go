package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
		panic(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	connectionString := os.Getenv("DB_CONNECTION_STRING")

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Automigrate the URL struct.
	db.AutoMigrate(&URL{},&User{})
	DB = db
}
