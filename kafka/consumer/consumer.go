package consumer

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func NewConsumer(topic string) {
	config := sarama.NewConfig()
	c, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}

	partitionC, err := c.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	consumed := 0
	
ConsumerLoop:
	for {
		select {
		case msg := <-partitionC.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
