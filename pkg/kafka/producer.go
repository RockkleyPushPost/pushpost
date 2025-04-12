package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(broker string, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
		//BatchTimeout: 1 * time.Second,
	}

	return &Producer{
		writer: writer,
	}
}

func (p *Producer) SendMessage(value []byte) error {
	if err := p.writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: value}); err != nil {
		log.Println("failed send to kafka: ", err)
		return err
	}
	return nil
}

func (p *Producer) Close() error {
	if err := p.writer.Close(); err != nil {
		log.Println("failed closing kafka writer: ", err)
		return err
	}
	return nil
}
