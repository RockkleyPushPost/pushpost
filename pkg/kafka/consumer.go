package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupID string) *Consumer {
	return &Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		}),
	}
}

func (c *Consumer) StartListening(handler func(message kafka.Message)) {
	go func() {
		log.Println("kafka consumer listening to the topic:", c.reader.Config().Topic)
		for {
			msg, err := c.reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("error reading from kafka:", err)
				continue
			}
			handler(msg)
		}
	}()
}

func (c *Consumer) Close() {
	if err := c.reader.Close(); err != nil {
		log.Println("error closing kafka reader:", err)
	}
	log.Println("kafka consumer closed")
}
