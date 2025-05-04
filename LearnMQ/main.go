package main

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  5 * time.Second,
		ConnectionTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	msgID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello world"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Published message with ID: %v", msgID)
}
