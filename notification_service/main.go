package main

import (
	"context"
	"log"
	"notification_service/database"
	"notification_service/kafka"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "notifications.db"
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	consumer := kafka.NewConsumer(brokers, "notification_group", database.SetupDB(dbPath))
	log.Printf("consuming notifications from %s", brokers)
	if err := consumer.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
