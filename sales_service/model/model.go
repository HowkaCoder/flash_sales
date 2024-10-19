package model

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey"`
	ProductID     uint      `gorm:"not null" json:"product_id"`
	DiscountPrice uint      `gorm:"not null" json:"discount_price"`
	StartTime     time.Time `gorm:"not null" json:"start_time"`
	EndTime       time.Time `gorm:"not null" json:"end_time"`
}
