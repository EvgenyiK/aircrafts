package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Aircraft struct {
	Code string `json:"aircraft_code" gorm:"primary_key"`
	Model string `json:"model"`
	Range int    `json:"range"`
}

var DB *gorm.DB

func Connect() *gorm.DB {
	err := godotenv.Load("/home/evgen/Development/aircrafts/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	

	dsn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(Aircraft{})
	DB = db
	return db
}
