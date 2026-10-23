package kafka

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

const topic = "notifications"

type Publisher struct {
	writer  *kafka.Writer
	service string
}

func NewPublisher(brokers, service string) *Publisher {
	if brokers == "" {
		return nil
	}
	return &Publisher{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(strings.Split(brokers, ",")...),
			Topic:                  topic,
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
		service: service,
	}
}

func (p *Publisher) Notify(event, message string) {
	if p == nil {
		return
	}

	payload, err := json.Marshal(map[string]string{
		"service_name": p.service,
		"event_type":   event,
		"message":      message,
	})
	if err != nil {
		log.Printf("kafka: marshal: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := p.writer.WriteMessages(ctx, kafka.Message{Key: []byte(event), Value: payload}); err != nil {
		log.Printf("kafka: publish %s: %v", event, err)
	}
}

func (p *Publisher) Close() error {
	if p == nil {
		return nil
	}
	return p.writer.Close()
}
