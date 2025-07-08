package main

import (
	"context"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	//创建客户端
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  5 * time.Second,
		ConnectionTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	//创建生产者
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my_topic",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()
	//发送消息
	msgID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello world"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Published message with ID: %v", msgID)
}
