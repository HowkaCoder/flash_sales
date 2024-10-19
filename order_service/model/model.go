package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null" json:"user_id"`
	ProductID uint   `gorm:"not null" json:"product_id"`
	Quantity  uint   `gorm:"not null" json:"quantity"`
	Status    string `gorm:"not null" json:"status"`
}
