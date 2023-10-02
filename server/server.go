package server

import (
	"project-client/api/http"
	"project-client/kafka/producer"
)

type Server interface {
	Start() error
	Stop()
}

type server struct {
	httpServer *http.GinServer
	producer   *producer.Producer
}

func (s *server) Start() error {
	producer, err := producer.NewProducer()
	if err != nil {
		return err
	}
	s.producer = producer
	s.httpServer = http.NewGinServer()
	s.httpServer.AddKafkaHandlers(producer)

	s.httpServer.Run()

	return nil
}

func (s *server) Stop() {

}

func NewServer() (Server, error) {
	return &server{}, nil
}
