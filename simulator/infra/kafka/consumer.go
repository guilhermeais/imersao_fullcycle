package kafka

import (
	"fmt"
	"log"
	"os"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	// MessageChannel é um canal de mensagens
	// que será utilizado para receber as mensagens e processá-las
	MessageChannel chan *cKafka.Message
}

func NewKafkaConsumer(messageChannel chan *cKafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChannel: messageChannel,
	}
}

func (kafkaConsumer *KafkaConsumer) Consume() {
	configMap := &cKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	consumer, err := cKafka.NewConsumer(configMap)

	if err != nil {
		log.Fatalf("error consuming kafka message: " + err.Error())
	}

	topics := []string{
		os.Getenv("KafkaReadTopic"),
	}

	consumer.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started!")

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			// Envia a mensagem para o canal de mensagens
			kafkaConsumer.MessageChannel <- msg
		}
	}
}
