package main

import (
	"context"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	//创建消费者
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  5 * time.Second,
		ConnectionTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	//订阅消息
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my_topic",
		SubscriptionName: "my-subscription",
		Type:             pulsar.Exclusive, // 订阅类型，可选 Shared / Failover / KeyShared
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	//如果收到消息，就消费打印出来
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("收到的消息为 : %s", string(msg.Payload()))

		// 确认消息
		consumer.Ack(msg)
	}
}
