package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"notification_service/model"
	"strings"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type Consumer struct {
	reader *kafka.Reader
	db     *gorm.DB
}

func NewConsumer(brokers, groupID string, db *gorm.DB) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: strings.Split(brokers, ","),
			GroupID: groupID,
			Topic:   "notifications",
		}),
		db: db,
	}
}

func (c *Consumer) Run(ctx context.Context) error {
	defer c.reader.Close()

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if errors.Is(err, context.Canceled) {
			return nil
		}
		if err != nil {
			return err
		}

		var n model.Notification
		if err := json.Unmarshal(msg.Value, &n); err != nil {
			log.Printf("skip message at offset %d: %v", msg.Offset, err)
			continue
		}
		n.Status = "received"

		if err := c.db.WithContext(ctx).Create(&n).Error; err != nil {
			log.Printf("save notification: %v", err)
			continue
		}
		log.Printf("%s: %s", n.ServiceName, n.Message)
	}
}
