package kafka

import (
	"log"
	"os"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaProducer() *cKafka.Producer {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	}

	producer, err := cKafka.NewProducer(configMap)

	if err != nil {
		log.Println("error creating kafka producer: %v", err)
	}

	return producer
}

func Publish(msg string, topic string, producer *cKafka.Producer) error {
	message := &cKafka.Message{
		TopicPartition: cKafka.TopicPartition{
			Topic:     &topic,
			Partition: cKafka.PartitionAny,
		},
		Value: []byte(msg),
	}

	err := producer.Produce(message, nil)

	if err != nil {
		return err
	}

	return nil
}
