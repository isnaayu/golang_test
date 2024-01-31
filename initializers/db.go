package initializers

import (
	"fmt"
	"os"
	"test_golang/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase(){
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}

func SyncDB(){
	// Migrate the schema
	DB.AutoMigrate(&models.Post{})
}