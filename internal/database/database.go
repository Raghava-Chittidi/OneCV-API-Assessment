package database

import (
	"log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatal("Failed to load in env file", err)
	// 	return nil, err
	// }

	db_url := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection string is invalid", err)
		return nil, err
	}

	DB = db
	log.Println("Connected to database!")
	return db, nil
}