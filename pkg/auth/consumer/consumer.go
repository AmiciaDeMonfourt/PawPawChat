package consumer

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	kafkaConsumer *kafka.Consumer
	MsgChannel    chan *kafka.Message
}

func New(brokers, groupID string) *Consumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		// "auto.offset.reset":  "earliest",
		"enable.auto.commit": true,
	}

	kafkaConsumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("failed to create kafka consumer: %v", err)
	}

	return &Consumer{
		kafkaConsumer: kafkaConsumer,
		MsgChannel:    make(chan *kafka.Message),
	}
}

func (c *Consumer) Subscribe(topics []string) error {
	return c.kafkaConsumer.Subscribe(topics[0], nil)
}

func (c *Consumer) Consume() {
	log.Printf("consumer receives messages")
	go func() {
		for {
			msg, err := c.kafkaConsumer.ReadMessage(-1)

			if err != nil {
				log.Printf("consumer error: %v \nmsg: (%v)\n", err, msg.Value)
				continue
			}

			log.Printf("accepted msg: %v", string(msg.Value))
			c.MsgChannel <- msg
		}
	}()
}
