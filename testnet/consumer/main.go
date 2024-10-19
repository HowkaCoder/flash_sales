package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Создаем конфигурацию консюмера
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Создаем консюмера
	consumer, err := sarama.NewConsumer([]string{"172.18.0.4:9093"}, config)
	if err != nil {
		log.Fatal("Error creating consumer:", err)
	}
	defer consumer.Close()

	// Подписываемся на топик
	partitionConsumer, err := consumer.ConsumePartition("your_topic_name", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Error consuming partition:", err)
	}
	defer partitionConsumer.Close()

	// Читаем сообщения
	for message := range partitionConsumer.Messages() {
		log.Printf("Received message: %s\n", string(message.Value))
	}
}
