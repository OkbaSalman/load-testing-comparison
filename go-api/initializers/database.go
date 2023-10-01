package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// sqlDB, errors := DB.DB()
	// sqlDB.SetMaxIdleConns(10)

	// if errors != nil {
	// 	log.Fatal("Failed to connenct to database")
	// }
	
	if err != nil {
		log.Fatal("Failed to connenct to database")
	}
}