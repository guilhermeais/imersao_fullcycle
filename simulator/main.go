package main

import (
	"fmt"
	"log"

	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"

	kafkaApplication "github.com/codeedu/imersaofsfc2-simulator/application/kafka"
	kafkaInfra "github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
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
	kafkaConsumer := kafkaInfra.NewKafkaConsumer(msgChan)
	go kafkaConsumer.Consume()

	for msg := range msgChan {
		go kafkaApplication.Produce(msg)
		fmt.Println(string(msg.Value))
	}
}
