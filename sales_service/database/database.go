package database

import (
	"log"
	"sales_service/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error with database initiating: %v", err)
	}

	db.AutoMigrate(&model.Sale{})

	return db
}
