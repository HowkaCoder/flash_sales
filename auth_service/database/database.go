package database

import (
	"auth_service/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&model.User{})

	return db
}
