package main

import (
	"fmt"
	"project-client/kafka/consumer"
	"project-client/kafka/producer"
	"time"

	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {
	
	fmt.Println("Server running")

	//Set up handlers
	// HTTP - GIN

	// RPC - GRPC

	//Initiallise Kafka Producer
	topic := "test-topic"

	go consumer.NewConsumer(topic)

	producer := producer.NewProducer()
	for i := 0; i < 100; i++ {
		producer.Send(fmt.Sprintf("TEST - %d", i), topic)
		time.Sleep(time.Second * 3)
	}
	
	fmt.Println("Server running")


	return nil
}
