package producer

import (
	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.AsyncProducer
}

func NewProducer() *Producer {

	config := sarama.NewConfig()
	asyncP, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	p := &Producer{}
	p.producer = asyncP

	return p
}

func (p *Producer) Send(data, topic string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.StringEncoder(data),
	}
	p.producer.Input() <- msg
}
