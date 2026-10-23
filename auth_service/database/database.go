package database

import (
	"auth_service/model"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func SetupDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path+"?_pragma=busy_timeout(5000)"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	return db
}
