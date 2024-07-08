package producer

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	kafkaProducer *kafka.Producer
}

func New(brokers string) *Producer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
		// "enable.idempotence": true,
		// "acks":    "all",
		// "retries": 5,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
	}

	return &Producer{kafkaProducer: producer}
}

func (p *Producer) Logs() {
	go func() {
		for e := range p.kafkaProducer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
					fmt.Printf("msg: %s", string(ev.Value))
				}
			}
		}
	}()
	// data, err := json.Marshal(msg)
	// if err != nil {
	// 	log.Printf("failed to marshal message to producer: %v", err)
	// 	return err
	// }

	// topic := "users"
	// message := &kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{
	// 		Topic:     &topic,
	// 		Partition: kafka.PartitionAny,
	// 	},
	// 	Value: data,
	// }

	// // channel used to receive delivery reports from producer in an async
	// deliveryChan := make(chan kafka.Event, 1)
	// defer close(deliveryChan)

	// if err = p.kafkaProducer.Produce(message, deliveryChan); err != nil {
	// 	log.Printf("failed to send message: %v", err)
	// 	return err
	// }

	// e := <-deliveryChan
	// m := e.(*kafka.Message)

	// if m.TopicPartition.Error != nil {
	// 	log.Printf("failed to send message: %v", err)
	// 	return err
	// }

	// log.Printf("message successfully sent in %v\n", m.TopicPartition)
	// log.Printf("message: %s\n", string(m.Value))
	// return nil
}

func (p *Producer) Send(msg interface{}) error {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("failed to marshal message to producer: %v", err)
		return err
	}

	topic := "users"
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
	}

	return p.kafkaProducer.Produce(message, nil)
}
