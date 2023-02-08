package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/codeedu/imersaofsfc2-simulator/application/route"
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	cKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// {"clientId": "1","id":"1"}
// {"clientId": "2","id":"2"}
func Produce(msg *cKafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route) // deserializa a mensagem recebida e armazena em route
	route.LoadPositions()

	positions, err := route.ExportJSONPositionsAsString()

	if err != nil {
		log.Println("Error to export positions as string", err.Error())
	}

	for _, position := range positions {
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
