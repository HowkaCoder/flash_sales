package main

import (
	"log"
	"notification_service/database"
	"notification_service/kafka"
)

func main() {
	db := database.SetupDB()

	//kafkaConsumer, err := kafka.NewKafkaConsumer("0.0.0.0:9092,0.0.0.0:9093,0.0.0.0:9094", "notification_group", db)
	kafkaConsumer, err := kafka.NewKafkaConsumer("kafka1:9092,kafka2:9093,kafka3:9094", "notification_group", db)

	if err != nil {
		log.Fatalf("Error initializing Kafka consumer: %v", err)
	}

	go kafkaConsumer.Start()

	select {} // Блокировка основного потока
}
