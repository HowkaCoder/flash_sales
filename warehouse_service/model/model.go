package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null" json:"name"`
	Desc     string `gorm:"not null" json:"desc"`
	Price    uint   `gorm:"not null" json:"price"`
	Quantity uint   `gorm:"not null" json:"quantity"`
}
