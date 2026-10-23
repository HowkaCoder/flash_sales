package database

import (
	"log"
	"warehouse_service/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func SetupDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path+"?_pragma=busy_timeout(5000)"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error with database st upping : %v !", err)
	}

	if err := db.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	return db

}
