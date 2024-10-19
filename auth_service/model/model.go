package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null" json:"name"`
	Address  string `gorm:"not null" json:"address"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
