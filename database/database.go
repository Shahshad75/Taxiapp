package database

import (
	"log"
	"os"
	"taxiapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dns := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.Driver{})
	DB.AutoMigrate(&models.VehicleDetails{})
	DB.AutoMigrate(&models.DriverDocument{})
}
