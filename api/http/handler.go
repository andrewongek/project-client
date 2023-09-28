package http

import "github.com/gin-gonic/gin"

func NewGinServer() *gin.Engine {
	r := gin.Default()

	return r
}