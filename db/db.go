package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitPostgresDb() {
	err = godotenv.Load(".env")

	if err != nil {
		log.Print("Error loading .env file", err)
	}

	var (
		host       = os.Getenv("DB_HOST")
		port       = os.Getenv("DB_PORT")
		dbUser     = os.Getenv("DB_USER")
		dbName     = os.Getenv("DB_NAME")
		dbPassword = os.Getenv("DB_PASSWORD")
	)

	dsm := fmt.Sprintf("host = %s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, dbUser, dbName, dbPassword)

	db, err = gorm.Open(postgres.Open(dsm), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Room{}, &Item{}, &Comment{})
}
