package main

import (
	"log"

	"github.com/IBM/sarama"
)

func main() {
	// Создаем конфигурацию продюсера
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Создаем продюсера
	producer, err := sarama.NewSyncProducer([]string{"172.18.0.4:9093"}, config)
	if err != nil {
		log.Fatal("Error creating producer:", err)
	}
	defer producer.Close()

	// Отправляем сообщение
	msg := &sarama.ProducerMessage{
		Topic: "your_topic_name",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Error sending message:", err)
	}
	log.Println("Message sent successfully")
}
