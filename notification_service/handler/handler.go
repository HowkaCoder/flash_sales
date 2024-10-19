package handler

import (
	"log"
	"notification_service/model"

	"gorm.io/gorm"
)

type NotificationHandler struct {
	DB *gorm.DB
}

func (h *NotificationHandler) GetNotifications() ([]model.Notification, error) {
	var notifications []model.Notification
	if err := h.DB.Find(&notifications).Error; err != nil {
		log.Printf("Error retrieving notifications: %v", err)
		return nil, err
	}
	return notifications, nil
}
