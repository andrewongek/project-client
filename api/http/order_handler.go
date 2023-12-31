package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"project-client/kafka/producer"

	"github.com/andrewongek/project-lib/proto/pb"
	"github.com/gin-gonic/gin"
)

func orderHandler(p *producer.Producer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, err.Error(), nil)
		}
		d := &pb.Order{}
		json.Unmarshal(data, d)

		p.Send(data, "test-topic")

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Notification sent successfully!",
		})
	}
}
