package model

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	ServiceName string `json:"service_name"`
	EventType   string `json:"event_type"`
	Message     string `json:"message"`
	Status      string `json:"status"`
}
