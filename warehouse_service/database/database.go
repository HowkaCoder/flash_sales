package database

import (
	"log"
	"warehouse_service/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error with database st upping : %v !", err)
	}

	db.AutoMigrate(&model.Product{})

	return db

}
