package producer

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	producer *kafka.Producer
}

func New() *Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		log.Fatal(err)
	}

	return &Producer{
		producer: producer,
	}
}

func (p *Producer) Send() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka:9092"})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	defer producer.Close()

	topic := "users"
	message := "Hello, Kafka!"

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		log.Fatalf("Failed to produce message: %s", err)
	}

	// producer.Flush(15 * 1000)
	fmt.Println("Message produced successfully")
}
