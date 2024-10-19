package kafka

import (
	"encoding/json"
	"log"
	"notification_service/model"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gorm.io/gorm"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer
	DB       *gorm.DB
}

func NewKafkaConsumer(brokers string, groupID string, db *gorm.DB) (*KafkaConsumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	consumer.SubscribeTopics([]string{"notifications"}, nil)
	return &KafkaConsumer{Consumer: consumer, DB: db}, nil
}

func (kc *KafkaConsumer) Start() {
	for {
		msg, err := kc.Consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var notification model.Notification
		if err := json.Unmarshal(msg.Value, &notification); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}
		log.Printf("Message : %v  and Service name : %v", notification.Message, notification.ServiceName)
		if err := kc.DB.Create(&notification).Error; err != nil {
			log.Printf("Error saving notification: %v", err)
		} else {
			log.Printf("Notification saved: %+v", notification)
		}
	}
}
