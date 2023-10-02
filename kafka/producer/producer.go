package producer

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.AsyncProducer
}

func NewProducer() (*Producer, error) {

	config := sarama.NewConfig()
	asyncP, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer", err)
	}

	p := &Producer{}
	p.producer = asyncP

	return p, nil
}

func (p *Producer) Send(data []byte, topic string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.StringEncoder(data),
	}
	p.producer.Input() <- msg
}
