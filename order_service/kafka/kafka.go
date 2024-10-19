package kafka

import (
	"fmt"
	"log"

	"context"

	"github.com/segmentio/kafka-go"
)

type ProducerConfig struct {
	Brokers []string
	Topic   string
}

func CreateProducer(cfg ProducerConfig) *kafka.Writer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  cfg.Brokers,
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	})

	return writer
}

func SendMessage(writer *kafka.Writer, key, value string) error {
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v\n", err)
		return err
	}
	fmt.Println("Сообщение успешно отправлено")
	return nil
}
func NotifyKafka(topic, key, message string) {
	producerConfig := ProducerConfig{
		Brokers: []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		Topic:   topic,
	}

	producer := CreateProducer(producerConfig)

	err := SendMessage(producer, key, message)
	if err != nil {
		fmt.Printf("Ошибка отправки сообщения: %v\n", err)
	}

	err = CloseProducer(producer)
	if err != nil {
		fmt.Printf("Ошибка закрытия продюсера: %v\n", err)
	}
}
func CloseProducer(writer *kafka.Writer) error {
	err := writer.Close()
	if err != nil {
		log.Printf("Ошибка закрытия продюсера: %v\n", err)
		return err
	}
	return nil
}
