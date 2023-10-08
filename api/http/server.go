package http

import (
	"net/http"
	"project-client/kafka/producer"

	"github.com/gin-gonic/gin"
)

type GinServer struct {
	router *gin.Engine
}

func NewGinServer() *GinServer {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server := &GinServer{}
	server.router = router
	return server
}

func (g *GinServer) AddKafkaHandlers(p *producer.Producer) *GinServer {
	g.router.POST("/make-order", orderHandler(p))
	return g
}

func (g *GinServer) Run() {
	g.router.Run()
}
