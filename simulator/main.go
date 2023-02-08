package main

import (
	"fmt"
	"log"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	msgChan := make(chan *cKafka.Message)
	kafkaConsumer := kafka.NewKafkaConsumer(msgChan)
	go kafkaConsumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}
}
